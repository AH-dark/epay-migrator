package main

import (
	"context"
	"os"

	"github.com/star-horizon/epay-database-mingrator/entry"
)

var ctx = context.Background()

func main() {
	if err := entry.Bootstrap(ctx); err != nil {
		panic(err)
	}

	os.Exit(0)
}
