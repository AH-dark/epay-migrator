package log

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"internal.log",
		fx.Provide(NewLogger),
	)
}
