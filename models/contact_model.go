package models

type Contact struct {
	ID uint `gorm:"primaryKey`
	Nome string
	Celular string
}