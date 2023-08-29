package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Batch struct {
	Batch    string          `gorm:"column:batch;primaryKey;not null"`
	AllMoney decimal.Decimal `gorm:"column:all_money;not null;type:decimal(10,2)"`
	Count    int             `gorm:"column:count;not null;default:0"`
	Time     time.Time       `gorm:"column:time;type:datetime"`
	Status   int             `gorm:"column:status;not null;default:0"`
}
