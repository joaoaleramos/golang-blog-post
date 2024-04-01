package service

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *entity.User) error {
	if err := s.userRepository.CreateUser(user); err != nil {
		return err
	}
	return nil

}

func (s *UserService) GetUserByID(userID uuid.UUID) (*entity.User, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err

	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
