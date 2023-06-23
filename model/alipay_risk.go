package model

import (
	"github.com/star-horizon/epay-database-mingrator/internal/conf"
	"time"
)

type AlipayRisk struct {
	ID           int       `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	Channel      uint      `gorm:"not null;column:channel" json:"channel"`
	Pid          string    `gorm:"not null;column:pid" json:"pid"`
	Smid         string    `gorm:"column:smid" json:"smid"`
	TradeNos     string    `gorm:"column:tradeNos" json:"tradeNos"`
	RiskType     string    `gorm:"column:risktype" json:"risktype"`
	RiskLevel    string    `gorm:"column:risklevel" json:"risklevel"`
	RiskDesc     string    `gorm:"column:riskdesc" json:"riskdesc"`
	ComplainTime string    `gorm:"column:complainTime" json:"complainTime"`
	ComplainText string    `gorm:"column:complainText" json:"complainText"`
	Date         time.Time `gorm:"column:date;type:datetime" json:"date"`
	Status       int       `gorm:"column:status;not null;default:0" json:"status"`
	ProcessCode  string    `gorm:"column:process_code" json:"process_code"`
}

func (AlipayRisk) TableName() string {
	return conf.DatabaseConfig.Prefix + "alipayrisk"
}
