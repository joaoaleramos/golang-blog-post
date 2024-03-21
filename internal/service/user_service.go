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

func (s *UserService) CreateUser(user *entity.User) error {
	if err := s.userRepository.Create(user); err != nil {
		return err
	}
	return nil

}

// TODO: finish method
// func (s *UserService) UpdateUser()
