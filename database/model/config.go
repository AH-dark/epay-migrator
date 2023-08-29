package model

type Config struct {
	Key string `gorm:"primaryKey;column:k" json:"k"`
	Val string `gorm:"column:v" json:"v"`
}
