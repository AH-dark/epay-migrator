package model

import "github.com/AH-dark/epay-migrator/internal/conf"

type Config struct {
	Key string `gorm:"primaryKey;column:k" json:"k"`
	Val string `gorm:"column:v" json:"v"`
}

func (Config) TableName() string {
	return conf.DatabaseConfig.Prefix + "config"
}
