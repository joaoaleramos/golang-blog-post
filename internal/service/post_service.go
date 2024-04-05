package service

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/domain/repository"

	"github.com/google/uuid"
)

type PostService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

func (s *PostService) CreatePost(post *entity.Post) error {
	err := s.postRepository.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPostByID(postID uuid.UUID) (*entity.Post, error) {
	post, err := s.postRepository.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	return post, nil

}

func (s *PostService) GetAllPosts() ([]*entity.Post, error) {
	posts, err := s.postRepository.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPostsByUserID(userID uuid.UUID) ([]*entity.Post, error) {
	posts, err := s.postRepository.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) UpdatePost(post *entity.Post) error {
	err := s.postRepository.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(postID uuid.UUID) error {
	err := s.postRepository.DeletePost(postID)
	if err != nil {
		return err
	}
	return nil
}
