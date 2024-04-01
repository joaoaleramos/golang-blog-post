package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type PostRepository interface {
	Create(user *entity.Post) (*entity.Post, error)
	GetPostByID(id uuid.UUID) (*entity.Post, error)
	GetPostsByUserID(user *entity.User) ([]*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	Update(Post *entity.Post) error
	Delete(id uuid.UUID) error
}
