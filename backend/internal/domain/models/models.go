package models

import (
	"time"
)

type User struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	TelegramID      uint      `json:"telegram_id" gorm:"unique"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username"`
	LanguageCode    string    `json:"language_code"`
	AllowsWriteToPm bool      `json:"allows_write_to_pm"`
	PhotoUrl        string    `json:"photo_url"`
	CreatedAt       time.Time `json:"created_at"`
}
