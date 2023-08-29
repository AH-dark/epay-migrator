package model

type Type struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	Name     string `gorm:"column:name;not null;index" json:"name"`
	Device   uint   `gorm:"column:device;not null;default:0;index" json:"device"`
	ShowName string `gorm:"column:showname;not null;default:''" json:"showname"`
	Status   int    `gorm:"column:status;not null;default:0" json:"status"`
}
