package models

import (
	"time"
)

type Food struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	FoodName        string    `json:"foodname"`
	Materials       string    `json:"materials"`
	EatPerson       int       `json:"eatperson"`
	Specification   string    `json:"specification"`
	GuessPrice      int       `json:"guessprice"`
	PreparationTime int       `json:"preparationtime"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ImageUrl        string    `json:"imageurl"`
	Image           string    `json:"image"`
}
