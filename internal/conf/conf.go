package conf

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	Version    = "1.0.0"
	InstanceID = uuid.NewString()
)

var validate = validator.New()

func LoadDotEnv(ctx context.Context) {
	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load env")
	}
}
