package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByUserID(id uuid.UUID) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id uuid.UUID) error
}
