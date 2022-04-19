package entity

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Title  string  `json:"title"`
}