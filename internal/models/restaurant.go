package models

import "time"

type Restaurant struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"not null"`
	Category  string     `json:"category" gorm:"not null"`
	Address   string     `json:"address" gorm:"not null"`
	ImageURL  string     `json:"image_url"`
	MenuItems []MenuItem `json:"menu_items,omitempty" gorm:"foreignKey:RestaurantID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type MenuItem struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RestaurantID uint      `json:"restaurant_id" gorm:"not null"`
	Name         string    `json:"name" gorm:"not null"`
	Price        float64   `json:"price" gorm:"not null"`
	Description  string    `json:"description"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	TotalPrice float64   `json:"total_price" gorm:"not null"`
	Status     string    `json:"status" gorm:"default:pending"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}