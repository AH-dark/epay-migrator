package conf

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type appConfig struct {
	SysKey  string `envconfig:"APP_SYSKEY"`
	CronKey string `envconfig:"APP_CRONKEY"`
}

var AppConfig appConfig

func InitAppConfig(ctx context.Context) error {
	prefix := "APP"

	if err := envconfig.Process(prefix, &AppConfig); err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("failed to load env %s", prefix)
		return err
	}

	logrus.WithContext(ctx).WithField("env", prefix).Debugf("loaded env %s: %+v", prefix, AppConfig)

	if err := validate.StructCtx(ctx, AppConfig); err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("failed to validate env %s", prefix)
		return err
	}

	return nil
}
