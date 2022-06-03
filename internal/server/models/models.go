package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	TgUserName   string
	TgId         int64
	TinkoffToken string
}
