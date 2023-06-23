package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(ctx context.Context) (*logrus.Logger, error) {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{})

	return logger, nil
}
