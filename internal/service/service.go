package service

import (
	"context"
	"task_manager/internal/models"
)

type UserService interface {
	Register(ctx context.Context, name, email, password string) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, error) // Возвращает JWT токен
}
