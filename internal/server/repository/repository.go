package repository

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	gdb *gorm.DB
}

func New(gdb *gorm.DB) *repository {
	return &repository{gdb: gdb}
}
