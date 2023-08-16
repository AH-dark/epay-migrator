package model

import (
	"github.com/AH-dark/epay-migrator/internal/conf"
	"time"
)

type Risk struct {
	ID      int       `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID     int       `gorm:"column:uid;not null;default:0;index" json:"uid"`
	Type    int       `gorm:"column:type;not null;default:0" json:"type"`
	Url     string    `gorm:"column:url" json:"url"`
	Content string    `gorm:"column:content" json:"content"`
	Date    time.Time `gorm:"column:date;type:datetime" json:"date"`
	Status  int       `gorm:"column:status;not null;default:0" json:"status"`
}

func (Risk) TableName() string {
	return conf.DatabaseConfig.Prefix + "risk"
}
