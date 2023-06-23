package model

import "github.com/star-horizon/epay-database-mingrator/internal/conf"

type RegCode struct {
	ID       int    `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	UID      int    `gorm:"column:uid;not null;default:0" json:"uid"`
	Type     int    `gorm:"column:type;not null;default:0;index" json:"type"`
	Code     string `gorm:"column:code;not null;index" json:"code"`
	To       string `gorm:"column:to;index" json:"to"`
	Time     int    `gorm:"column:time;not null" json:"time"`
	IP       string `gorm:"column:ip" json:"ip"`
	Status   int    `gorm:"column:status;not null;default:0" json:"status"`
	ErrCount int    `gorm:"column:errcount;not null;default:0" json:"errcount"`
}

func (RegCode) TableName() string {
	return conf.DatabaseConfig.Prefix + "regcode"
}
