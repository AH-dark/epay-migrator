package model

import "github.com/AH-dark/epay-migrator/internal/conf"

type Weixin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	Type      uint   `gorm:"column:type;not null;default:0" json:"type"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Status    int    `gorm:"column:status;not null;default:0" json:"status"`
	AppID     string `gorm:"column:appid" json:"appid"`
	AppSecret string `gorm:"column:appsecret" json:"appsecret"`
}

func (Weixin) TableName() string {
	return conf.DatabaseConfig.Prefix + "weixin"
}
