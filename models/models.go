package models

import "time"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	ImageURL    string
	Category    string
	Stock       int
}

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Address    string
	Phone      string
	CreatedAt  time.Time
	OrderItems []OrderItem
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
