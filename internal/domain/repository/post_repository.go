package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
)

type PostRepository interface {
	CreatePost(post *entity.Post) error
	GetPostByID(postID uuid.UUID) (*entity.Post, error)
	GetPostsByUserID(userID uuid.UUID) ([]*entity.Post, error)
	GetAllPosts() ([]*entity.Post, error)
	UpdatePost(post *entity.Post) error
	DeletePost(postID uuid.UUID) error
}
