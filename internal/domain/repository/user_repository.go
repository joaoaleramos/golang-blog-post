package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(user_id uuid.UUID) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(user_id uuid.UUID) error
}
