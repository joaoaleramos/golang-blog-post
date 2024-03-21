package repository

import (
	"blog-post/internal/domain/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLUserRepository struct {
	DB *sqlx.DB
}

func NewSQLUserRepository(db *sqlx.DB) *SQLUserRepository {
	return &SQLUserRepository{DB: db}
}

func (r *SQLUserRepository) Create(user *entity.User) error {
	_, err := r.DB.Exec(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLUserRepository) GetByUserID(id uuid.UUID) (*entity.User, error) {
	// Implementation of GetByUserID method
	return &entity.User{}, nil
}

func (r *SQLUserRepository) GetAll() ([]*entity.User, error) {
	// Implementation of GetAll method
	return []*entity.User{}, nil
}

func (r *SQLUserRepository) Update(user *entity.User) error {
	// Implementation of Update method
	return nil
}

func (r *SQLUserRepository) Delete(id uuid.UUID) error {
	// Implementation of the Delete method goes here
	return nil
}
