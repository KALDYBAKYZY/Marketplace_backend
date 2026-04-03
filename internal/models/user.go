package models

type User struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
