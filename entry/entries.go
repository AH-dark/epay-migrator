package entry

import (
	"github.com/star-horizon/epay-database-mingrator/internal/log"
	"go.uber.org/fx"

	"github.com/star-horizon/epay-database-mingrator/internal/conf"
	"github.com/star-horizon/epay-database-mingrator/internal/infra"
)

func Entries() []fx.Option {
	return []fx.Option{
		conf.Module(),
		log.Module(),
		fx.WithLogger(log.FxLogger),
		infra.Module(),
	}
}
