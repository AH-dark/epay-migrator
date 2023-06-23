package entry

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func Bootstrap(ctx context.Context) error {
	opts := []fx.Option{
		fx.Supply(fx.Annotate(ctx, fx.As(new(context.Context)))),
		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
	}
	opts = append(opts, Entries()...)
	opts = append(opts, fx.Invoke(runMigrate))

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
