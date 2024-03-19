package service

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(name string, email string, password string) (entity.User, error) {
	user := entity.NewUser(
		name,
		email,
		password,
	)
	if err := s.userRepository.Create(user); err != nil {
		return entity.User{}, err
	}
	return *user, nil

}

// TODO: finish method
// func (s *UserService) UpdateUser()
