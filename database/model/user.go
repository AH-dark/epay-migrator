package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	UID          uint            `gorm:"primaryKey;autoIncrement;not null;column:uid" json:"uid"`
	GID          uint            `gorm:"column:gid;not null;default:0" json:"gid"`
	UpID         uint            `gorm:"column:upid;not null;default:0" json:"upid"`
	Key          string          `gorm:"column:key;not null" json:"key"`
	Password     string          `gorm:"column:pwd" json:"pwd"`
	Account      string          `gorm:"column:account" json:"account"`
	Username     string          `gorm:"column:username" json:"username"`
	CodeName     string          `gorm:"column:codename" json:"codename"`
	SettleID     int             `gorm:"column:settle_id;not null;default:1" json:"settle_id"`
	AlipayUid    string          `gorm:"column:alipay_uid" json:"alipay_uid"`
	QQUid        string          `gorm:"column:qq_uid" json:"qq_uid"`
	WxUid        string          `gorm:"column:wx_uid" json:"wx_uid"`
	Money        decimal.Decimal `gorm:"column:money;not null;default:0.00" json:"money"`
	Email        string          `gorm:"column:email;index" json:"email"`
	Phone        string          `gorm:"column:phone;index" json:"phone"`
	QQ           string          `gorm:"column:qq" json:"qq"`
	Url          string          `gorm:"column:url" json:"url"`
	Cert         int             `gorm:"column:cert;not null;default:0" json:"cert"`
	CertType     int             `gorm:"column:certtype;not null;default:0" json:"certtype"`
	CertNo       string          `gorm:"column:certno" json:"certno"`
	CertName     string          `gorm:"column:certname" json:"certname"`
	CertTime     time.Time       `gorm:"column:certtime;type:datetime" json:"certtime"`
	CertToken    string          `gorm:"column:certtoken" json:"certtoken"`
	CertCorpNo   string          `gorm:"column:certcorpno" json:"certcorpno"`
	CertCorpName string          `gorm:"column:certcorpname" json:"certcorpname"`
	AddTime      time.Time       `gorm:"column:addtime;type:datetime" json:"addtime"`
	LastTime     time.Time       `gorm:"column:lasttime;type:datetime" json:"lasttime"`
	EndTime      time.Time       `gorm:"column:endtime;type:datetime" json:"endtime"`
	Level        int             `gorm:"column:level;not null;default:1" json:"level"`
	Pay          int             `gorm:"column:pay;not null;default:1" json:"pay"`
	Settle       int             `gorm:"column:settle;not null;default:1" json:"settle"`
	KeyLogin     int             `gorm:"column:keylogin;not null;default:1" json:"keylogin"`
	Apply        int             `gorm:"column:apply;not null;default:0" json:"apply"`
	Mode         int             `gorm:"column:mode;not null;default:0" json:"mode"`
	Status       int             `gorm:"column:status;not null;default:0" json:"status"`
	Refund       int             `gorm:"column:refund;not null;default:0" json:"refund"`
	ChannelInfo  string          `gorm:"column:channelinfo" json:"channelinfo"`
	OrderName    string          `gorm:"column:ordername" json:"ordername"`
}
