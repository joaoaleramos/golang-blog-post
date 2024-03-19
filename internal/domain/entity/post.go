package entity

import "github.com/google/uuid"

type Post struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
}

func NewPost(userID uuid.UUID, title string, body string) *Post {
	return &Post{
		ID:     uuid.New(),
		UserID: userID,
		Title:  title,
		Body:   body,
	}
}
