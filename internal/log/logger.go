package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(ctx context.Context, debug bool) (*logrus.Logger, error) {
	logger := logrus.StandardLogger()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{})

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger, nil
}
