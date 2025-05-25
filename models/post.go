package models

import "time"

type Post struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"My Post"`
	Content   string    `json:"content" example:"Content goes here"`
	Comments  []Comment `json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"comments" gorm:"foreignKey:PostID"`
}
