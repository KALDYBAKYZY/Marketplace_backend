package models

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"unique"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
