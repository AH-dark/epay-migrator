package entry

import (
	"go.uber.org/fx"

	"github.com/AH-dark/epay-migrator/internal/conf"
	"github.com/AH-dark/epay-migrator/internal/infra"
	"github.com/AH-dark/epay-migrator/internal/log"
)

func Entries() []fx.Option {
	return []fx.Option{
		conf.Module(),
		log.Module(),
		fx.WithLogger(log.FxLogger),
		infra.Module(),
	}
}
