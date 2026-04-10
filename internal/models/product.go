package models

type Product struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"unique"`
	Price      float64
	Stock      int
	CategoryID uint `json:"category_id"`
}
