package model

type Plugin struct {
	Name     string `gorm:"primaryKey;column:name" json:"name"`
	ShowName string `gorm:"column:showname" json:"showname"`
	Author   string `gorm:"column:author" json:"author"`
	Link     string `gorm:"column:link" json:"link"`
	Types    string `gorm:"column:types" json:"types"`
}
