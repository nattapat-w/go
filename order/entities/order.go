package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID    uint
	Total float64
}
