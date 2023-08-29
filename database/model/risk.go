package model

import (
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
