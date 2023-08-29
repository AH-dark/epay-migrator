package main

import (
	"context"
	"github.com/AH-dark/epay-migrator/bootstrap"
	"os"

	"github.com/AH-dark/epay-migrator/internal/log"
)

func main() {
	ctx := context.Background()

	if err := bootstrap.GetApp().RunContext(ctx, os.Args); err != nil {
		log.Log(ctx).WithError(err).Fatal("failed to run command")
		os.Exit(1)
		return
	}
}
