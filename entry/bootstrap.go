package entry

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/star-horizon/epay-database-mingrator/actions"
)

type BootstrapParams struct {
	Action  string `name:"action"`
	IsDebug bool   `name:"debug"`
}

func Bootstrap(ctx context.Context, params BootstrapParams) error {
	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			params,
			params.Action,
			params.IsDebug,
		),
		fx.StartTimeout(time.Minute * 5),
		fx.StopTimeout(time.Minute * 5),
		fx.RecoverFromPanics(),
	}
	opts = append(opts, Entries()...)
	opts = append(opts, fx.Invoke(func(ctx context.Context, action string, debug bool) {
		logrus.WithContext(ctx).Infof("action: %s", action)
		logrus.WithContext(ctx).Infof("debug: %t", debug)
	}))

	switch params.Action {
	case "migrate":
		opts = append(opts, fx.Invoke(actions.RunMigrate))
	case "":
		logrus.WithContext(ctx).Error("action is required")
		return errors.New("action is required")
	default:
		logrus.WithContext(ctx).Errorf("unknown action: %s", params.Action)
		return fmt.Errorf("unknown action: %s", params.Action)
	}

	app := fx.New(opts...)

	if err := app.Start(ctx); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to start application")
		return err
	}

	logrus.WithContext(ctx).Info("application started")
	app.Wait()

	// Return the error if the application stopped with an error
	if err := app.Err(); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("application stopped with error")
		return err
	}

	logrus.WithContext(ctx).Info("application stopped")

	return nil
}
