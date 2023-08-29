package model

import (
	"github.com/shopspring/decimal"
)

type Group struct {
	GID        uint            `gorm:"primaryKey;autoIncrement;not null;column:gid" json:"gid"`
	Name       string          `gorm:"column:name;not null;default:''" json:"name"`
	Info       string          `gorm:"column:info;default:''" json:"info"`
	IsBuy      bool            `gorm:"column:isbuy;default:0" json:"is_buy"`
	Price      decimal.Decimal `gorm:"column:price;type:decimal(10,2);default:0.00" json:"price"`
	Sort       int             `gorm:"column:sort;default:0;not null" json:"sort"`
	Expire     int             `gorm:"column:expire;default:0;not null" json:"expire"`
	SettleOpen int             `gorm:"column:settle_open;default:0;not null" json:"settle_open"`
	SettleType int             `gorm:"column:settle_type;default:0;not null" json:"settle_type"`
	Settings   string          `gorm:"column:settings;default:''" json:"settings"`
}
