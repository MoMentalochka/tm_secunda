package user

import (
	"context"
	"task_manager/internal/models"
	"task_manager/internal/repository"
)

type tmUserService struct {
	userRepository repository.UserRepository
}

func NewTmUserService(userRepository repository.UserRepository) *tmUserService {
	return &tmUserService{
		userRepository,
	}
}
func (s *tmUserService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	return nil, nil
}
func (s *tmUserService) Login(ctx context.Context, email, password string) (string, error) {
	return "123", nil
}
