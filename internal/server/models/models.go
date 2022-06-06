package models

import "gorm.io/gorm"

type Id struct {
	Id int64
}

type TgId struct {
	TgId int64
}

type User struct {
	gorm.Model
	Id           int64
	Name         string
	TgUserName   string
	TgId         int64
	TinkoffToken string
}
