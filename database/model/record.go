package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Record struct {
	ID       int             `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID      int             `gorm:"column:uid;not null;index" json:"uid"`
	Action   int             `gorm:"column:action;not null" json:"action"`
	Money    decimal.Decimal `gorm:"column:money;type:decimal(10,2)" json:"money"`
	OldMoney decimal.Decimal `gorm:"column:oldmoney;type:decimal(10,2)" json:"oldmoney"`
	NewMoney decimal.Decimal `gorm:"column:newmoney;type:decimal(10,2)" json:"newmoney"`
	Type     string          `gorm:"column:type" json:"type"`
	TradeNo  string          `gorm:"column:trade_no;index" json:"trade_no"`
	Date     time.Time       `gorm:"column:date;type:datetime;not null" json:"date"`
}
