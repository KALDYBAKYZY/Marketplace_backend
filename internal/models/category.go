package models

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"unique"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
