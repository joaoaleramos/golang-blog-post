package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(userID uuid.UUID) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(userID uuid.UUID) error
}
