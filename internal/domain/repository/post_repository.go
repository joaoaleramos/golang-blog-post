package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type PostRepository interface {
	Create(user *entity.Post) (*entity.Post, error)
	GetPostByID(postID uuid.UUID) (*entity.Post, error)
	GetPostsByUserID(userID uuid.UUID) ([]*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(postID uuid.UUID) error
}
