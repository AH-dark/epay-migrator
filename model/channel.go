package model

import (
	"github.com/AH-dark/epay-migrator/internal/conf"
	"github.com/shopspring/decimal"
)

type Channel struct {
	ID            uint            `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Mode          int             `gorm:"column:mode;default:0" json:"mode"`
	Type          uint            `gorm:"column:type;not null" json:"type"`
	Plugin        string          `gorm:"column:plugin;not null" json:"plugin"`
	Name          string          `gorm:"column:name;not null" json:"name"`
	Rate          decimal.Decimal `gorm:"column:rate;not null;type:decimal(5,2);default:100.00" json:"rate"`
	Status        int             `gorm:"column:status;not null;default:0" json:"status"`
	AppID         string          `gorm:"column:appid" json:"appid"`
	AppKey        string          `gorm:"column:appkey" json:"appkey"`
	AppSecret     string          `gorm:"column:appsecret" json:"appsecret"`
	AppUrl        string          `gorm:"column:appurl" json:"appurl"`
	AppMerchantID string          `gorm:"column:appmchid" json:"appmchid"`
	AppType       string          `gorm:"column:apptype" json:"apptype"`
	DayTop        int             `gorm:"column:daytop;default:0" json:"daytop"`
	DayStatus     int             `gorm:"column:daystatus;default:0" json:"daystatus"`
	PayMin        string          `gorm:"column:paymin;default:0" json:"paymin"`
	PayMax        string          `gorm:"column:paymax;default:0" json:"paymax"`
	AppWxMp       int             `gorm:"column:appwxmp" json:"appwxmp"`
	AppWxA        int             `gorm:"column:appwxa" json:"appwxa"`
	AppSwitch     int             `gorm:"column:appswitch" json:"appswitch"`
}

func (Channel) TableName() string {
	return conf.DatabaseConfig.Prefix + "channel"
}
