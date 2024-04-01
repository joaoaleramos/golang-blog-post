package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLPostRepository struct {
	DB *sqlx.DB
}

func NewSQLPostRepository(db *sqlx.DB) *SQLPostRepository {
	return &SQLPostRepository{DB: db}
}

func (r *SQLPostRepository) CreatePost(post *entity.Post) error {
	return nil
}

func (r *SQLPostRepository) GetPostByID(id uuid.UUID) (*entity.Post, error) {
	// Define a Post struct to store the result
	var post entity.Post

	return &post, nil
}

func (r *SQLPostRepository) GetAllPosts() ([]*entity.Post, error) {
	// Execute the query
	var posts []*entity.Post

	return posts, nil
}

func (r *SQLPostRepository) UpdatePost(Post *entity.Post) error {
	// Implementation of Update method
	return nil
}

func (r *SQLPostRepository) DeletePost(id uuid.UUID) error {
	// Implementation of the Delete method goes here
	return nil
}
