package models

type Product struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	Name       string
	Price      float64
	Stock      int
	CategoryID uint
}
