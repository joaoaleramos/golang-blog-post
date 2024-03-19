package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}
