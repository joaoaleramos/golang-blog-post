package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPost(userID uuid.UUID, title string, body string, createdAt time.Time, updatedAt time.Time) *Post {
	return &Post{
		ID:        uuid.New(),
		UserID:    userID,
		Title:     title,
		Body:      body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
