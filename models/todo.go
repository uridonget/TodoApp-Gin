package models

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
}