package model

import (
	"github.com/star-horizon/epay-database-mingrator/internal/conf"
	"time"
)

type Anounce struct {
	ID      int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Content string    `gorm:"column:content" json:"content"`
	Color   string    `gorm:"column:color" json:"color"`
	Sort    int       `gorm:"column:sort;not null;default:1" json:"sort"`
	AddTime time.Time `gorm:"column:addtime;type:datetime" json:"add_time"`
	Status  int       `gorm:"column:status;not null;default:1" json:"status"`
}

func (Anounce) TableName() string {
	return conf.DatabaseConfig.Prefix + "anounce"
}
