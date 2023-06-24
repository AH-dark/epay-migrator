package main

import (
	"context"
	"flag"
	"os"

	"github.com/star-horizon/epay-database-mingrator/entry"
)

var ctx = context.Background()

var (
	isDebug bool
	isHelp  bool
	action  string
)

func init() {
	flag.BoolVar(&isDebug, "debug", false, "enable debug mode")
	flag.BoolVar(&isHelp, "help", false, "show help")
	flag.StringVar(&action, "action", "", "action to run, available actions: migrate")
	flag.Parse()
}

func main() {
	if isHelp || action == "" {
		flag.PrintDefaults()
		return
	}

	if err := entry.Bootstrap(ctx, entry.BootstrapParams{
		Action:  action,
		IsDebug: isDebug,
	}); err != nil {
		panic(err)
	}

	os.Exit(0)
}
