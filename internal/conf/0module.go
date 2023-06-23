package conf

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("conf",
		fx.Invoke(LoadDotEnv),
		fx.Invoke(InitAppConfig),
		fx.Invoke(InitDBConfig),
	)
}
