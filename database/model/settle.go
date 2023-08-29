package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Settle struct {
	ID             int             `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID            int             `gorm:"column:uid;not null;default:0;index" json:"uid"`
	Batch          string          `gorm:"column:batch;index" json:"batch"`
	Auto           int             `gorm:"column:auto;not null;default:1" json:"auto"`
	Type           int             `gorm:"column:type;not null;default:1" json:"type"`
	Account        string          `gorm:"column:account;not null" json:"account"`
	Username       string          `gorm:"column:username;not null" json:"username"`
	Money          decimal.Decimal `gorm:"column:money;not null" json:"money"`
	RealMoney      decimal.Decimal `gorm:"column:realmoney;not null" json:"realmoney"`
	AddTime        time.Time       `gorm:"column:addtime;type:datetime" json:"addtime"`
	EndTime        time.Time       `gorm:"column:endtime;type:datetime" json:"endtime"`
	Status         int             `gorm:"column:status;not null;default:0" json:"status"`
	TransferStatus int             `gorm:"column:transfer_status;not null;default:0" json:"transfer_status"`
	TransferResult string          `gorm:"column:transfer_result" json:"transfer_result"`
	TransferDate   time.Time       `gorm:"column:transfer_date;type:datetime" json:"transfer_date"`
	Result         string          `gorm:"column:result" json:"result"`
}
