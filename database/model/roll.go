package model

type Roll struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;not null;column:id" json:"id"`
	Type   uint   `gorm:"column:type;not null;default:0" json:"type"`
	Name   string `gorm:"column:name;not null;default:''" json:"name"`
	Kind   uint   `gorm:"column:kind;not null;default:0" json:"kind"`
	Info   string `gorm:"column:info" json:"info"`
	Status int    `gorm:"column:status;not null;default:0" json:"status"`
	Index  int    `gorm:"column:index;not null;default:0" json:"index"`
}
