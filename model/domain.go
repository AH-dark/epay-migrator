package model

import (
	"github.com/star-horizon/epay-database-mingrator/internal/conf"
	"time"
)

type Domain struct {
	ID      uint      `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID     int       `gorm:"column:uid;not null;default:0" json:"uid"`
	Domain  string    `gorm:"column:domain;not null;default:''" json:"domain"`
	Status  int       `gorm:"column:status;not null;default:0" json:"status"`
	AddTime time.Time `gorm:"column:addtime;type:datetime" json:"addtime"`
	EndTime time.Time `gorm:"column:endtime;type:datetime" json:"endtime"`
}

func (Domain) TableName() string {
	return conf.DatabaseConfig.Prefix + "domain"
}
