package user

import (
	"context"
	"database/sql"
	"errors"
	"task_manager/internal/models"

	"github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

// Create вставляет нового пользователя в БД
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password_hash, name, created_at) VALUES (?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, user.Email, user.PasswordHash, user.Name, user.CreatedAt)
	if err != nil {
		// Проверяем ошибку на дубликат Email (MySQL Error 1062)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return models.ErrEmailAlreadyExists // Определим эту ошибку в domain
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

// GetByEmail ищет пользователя по email для процесса Login
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash, name, created_at FROM users WHERE email = ?`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

// GetByID понадобится для проверки существования пользователей при инвайтах
func (r *userRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, email, password_hash, name, created_at FROM users WHERE id = ?`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
