package repository

import (
	"blog-post/internal/domain/entity"
	"log/slog"

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
	_, err := r.DB.Exec(`INSERT INTO posts (id, user_id, title, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) `, post.ID, post.UserID, post.Title, post.Body, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLPostRepository) GetPostByID(postID uuid.UUID) (*entity.Post, error) {
	// Define a Post struct to store the result
	var post entity.Post

	// Execute a query to find a post by ID
	err := r.DB.Get(&post, "SELECT id FROM posts WHERE id = $1", postID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *SQLPostRepository) GetPostsByUserID(userID uuid.UUID) ([]*entity.Post, error) {
	var posts []*entity.Post

	// Execute the query to fetch posts by user ID
	rows, err := r.DB.Queryx("SELECT * FROM posts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan each into a Post struct
	for rows.Next() {
		var post entity.Post
		if err := rows.StructScan(&post); err != nil {
			slog.Error("Error scanning post")
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *SQLPostRepository) GetAllPosts() ([]*entity.Post, error) {
	// Execute the query
	var posts []*entity.Post
	var post *entity.Post
	rows, err := r.DB.Queryx("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&post); err != nil {
			slog.Error("Error scanning post")
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *SQLPostRepository) UpdatePost(post *entity.Post) error {
	// Execute the query to update the post
	_, err := r.DB.Exec(`UPDATE posts SET title = $1, body = $2, updated_at = $3 WHERE id = $4`, post.Title, post.Body, post.UpdatedAt, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLPostRepository) DeletePost(postID uuid.UUID) error {
	// Execute the query to delete the post
	_, err := r.DB.Exec("DELETE FROM posts WHERE id = $1", postID)
	if err != nil {
		return err
	}
	return nil
}
