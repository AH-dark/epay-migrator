package model

import (
	"github.com/AH-dark/epay-migrator/internal/conf"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	TradeNo     string          `gorm:"primaryKey;column:trade_no;type:char(19);not null" json:"trade_no"`
	OutTradeNo  string          `gorm:"column:out_trade_no;not null;index" json:"out_trade_no"`
	ApiTradeNo  string          `gorm:"column:api_trade_no;index" json:"api_trade_no"`
	UID         uint            `gorm:"column:uid;not null;index" json:"uid"`
	TID         uint            `gorm:"column:tid;not null;default:0" json:"tid"`
	Type        uint            `gorm:"column:type;not null;" json:"type"`
	Channel     uint            `gorm:"column:channel;not null" json:"channel"`
	Name        string          `gorm:"column:name;not null" json:"name"`
	Money       decimal.Decimal `gorm:"column:money;type:decimal(10,2);not null" json:"money"`
	RealMoney   decimal.Decimal `gorm:"column:realmoney;type:decimal(10,2)" json:"real_money"`
	GetMoney    decimal.Decimal `gorm:"column:getmoney;type:decimal(10,2)" json:"get_money"`
	NotifyUrl   string          `gorm:"column:notify_url;type:varchar(255)" json:"notify_url"`
	ReturnUrl   string          `gorm:"column:return_url;type:varchar(255)" json:"return_url"`
	Param       string          `gorm:"column:param" json:"param"`
	AddTime     time.Time       `gorm:"column:addtime;type:datetime" json:"addtime"`
	EndTime     time.Time       `gorm:"column:endtime;type:datetime" json:"endtime"`
	Date        time.Time       `gorm:"column:date;type:date;index" json:"date"`
	Domain      string          `gorm:"column:domain" json:"domain"`
	Domain2     string          `gorm:"column:domain2" json:"domain2"`
	IP          string          `gorm:"column:ip" json:"ip"`
	Buyer       string          `gorm:"column:buyer" json:"buyer"`
	Status      int             `gorm:"column:status;not null;default:0" json:"status"`
	Notify      int             `gorm:"column:notify;not null;default:0" json:"notify"`
	NotifyTime  time.Time       `gorm:"column:notifytime;type:datetime" json:"notifytime"`
	Invite      uint            `gorm:"column:invite;not null;default:0;index" json:"invite"`
	InviteMoney decimal.Decimal `gorm:"column:invite_money;type:decimal(10,2)" json:"invite_money"`
}

func (Order) TableName() string {
	return conf.DatabaseConfig.Prefix + "order"
}
