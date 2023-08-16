package model

import (
	"github.com/AH-dark/epay-migrator/internal/conf"
	"time"
)

type Log struct {
	ID   int       `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID  int       `gorm:"column:uid;not null;default:0" json:"uid"`
	Type string    `gorm:"column:type" json:"type"`
	Date time.Time `gorm:"column:date;type:datetime" json:"date"`
	IP   string    `gorm:"column:ip" json:"ip"`
	City string    `gorm:"column:city" json:"city"`
	Data string    `gorm:"column:data" json:"data"`
}

func (Log) TableName() string {
	return conf.DatabaseConfig.Prefix + "log"
}
