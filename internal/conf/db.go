package conf

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type databaseConfig struct {
	Driver   string `default:"mysql"`
	Host     string `default:"localhost" validate:"hostname"`
	Port     int    `default:"3306" validate:"min=1,max=65535"`
	Name     string `default:"pay"`
	Username string `default:"root"`
	Password string `default:"root"`
	Prefix   string `default:"pre_"`
	SSLMode  string `default:"disable" envconfig:"DB_SSL_MODE"`
}

var DatabaseConfig databaseConfig

func InitDBConfig() {
	prefix := "DB"
	ctx := context.Background()

	if err := envconfig.Process(prefix, &DatabaseConfig); err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("failed to load env %s", prefix)
		panic(err)
	}

	logrus.WithContext(ctx).WithField("env", prefix).Debugf("loaded env %s: %+v", prefix, DatabaseConfig)

	if err := validate.StructCtx(ctx, DatabaseConfig); err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("failed to validate env %s", prefix)
		panic(err)
	}
}
