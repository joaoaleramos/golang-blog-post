package repository

import (
	"blog-post/internal/domain/entity"

	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLUserRepository struct {
	DB *sqlx.DB
}

func NewSQLUserRepository(db *sqlx.DB) *SQLUserRepository {
	return &SQLUserRepository{DB: db}
}

func (r *SQLUserRepository) CreateUser(user *entity.User) error {
	_, err := r.DB.Exec(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLUserRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	// Define a user struct to store the result
	var user entity.User

	// Execute the query to fetch the user by ID
	err := r.DB.Get(&user, "SELECT id, name, email, password FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SQLUserRepository) GetAllUsers() ([]*entity.User, error) {
	// Execute the query
	var users []*entity.User
	rows, err := r.DB.Queryx("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.StructScan(&user); err != nil {
			log.Fatal("Error scanning user")
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *SQLUserRepository) UpdateUser(user *entity.User) error {
	// Implementation of Update method
	return nil
}

func (r *SQLUserRepository) DeleteUser(id uuid.UUID) error {
	// Implementation of the Delete method goes here
	return nil
}
