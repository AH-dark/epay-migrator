package infra

import (
	"context"
	"fmt"

	gormlogrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/AH-dark/epay-migrator/internal/conf"
)

func NewDatabase(ctx context.Context) (*gorm.DB, error) {
	var dialect gorm.Dialector
	switch conf.DatabaseConfig.Driver {
	case "mysql", "mariadb":
		dialect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DatabaseConfig.Username,
			conf.DatabaseConfig.Password,
			conf.DatabaseConfig.Host,
			conf.DatabaseConfig.Port,
			conf.DatabaseConfig.Name,
		))
	case "postgres", "postgresql":
		dialect = postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			conf.DatabaseConfig.Host,
			conf.DatabaseConfig.Port,
			conf.DatabaseConfig.Username,
			conf.DatabaseConfig.Password,
			conf.DatabaseConfig.Name,
			conf.DatabaseConfig.SSLMode,
		))
	default:
		err := fmt.Errorf("unsupported database driver: %s", conf.DatabaseConfig.Driver)
		logrus.WithContext(ctx).WithError(err).Error("select database driver failed")
		return nil, err
	}

	db, err := gorm.Open(dialect, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.DatabaseConfig.Prefix,
			SingularTable: true,
		},
		Logger: gormlogrus.New(),
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("open database failed")
		return nil, err
	}

	if conf.DatabaseConfig.Driver == "mysql" || conf.DatabaseConfig.Driver == "mariadb" {
		db.Set("gorm:table_options", "DEFAULT CHARSET = utf8 ENGINE = InnoDB")
	}

	return db, nil
}
