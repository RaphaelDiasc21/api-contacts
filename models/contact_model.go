package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID uint `gorm:"primaryKey`
	Nome string
	Celular string
}