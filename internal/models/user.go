package models

import (
	"context"
	"time"
)

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	CreatedAt    time.Time `db:"created_at"`
}

// AuthUsecase описывает бизнес-логику авторизации
type AuthUsecase interface {
	Register(ctx context.Context, name, email, password string) (*User, error)
	Login(ctx context.Context, email, password string) (string, error) // Возвращает JWT токен
}

// UserRepository описывает контракт работы с БД для пользователей
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
}
