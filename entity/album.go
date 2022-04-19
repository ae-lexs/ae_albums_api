package entity

type Album struct {
	// gorm.Model
	Artist string  `json:"artist"`
	Price  float64 `json:"price,string"`
	Title  string  `json:"title"`
}
