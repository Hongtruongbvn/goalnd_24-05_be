package models

import "time"

type Comment struct {
	ID        uint      `json:"id" example:"1"`
	Content   string    `json:"content" example:"Nice post!"`
	PostID    uint      `json:"post_id" example:"10"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"post_id"`
}
