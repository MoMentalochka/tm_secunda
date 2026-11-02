package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

// Claims структура для JWT полезной нагрузки
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTAuthMiddleware валидирует токен и кладет userID в контекст
func JWTAuthMiddleware(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error": "missing authorization header"}`, http.StatusUnauthorized)
				return
			}

			// Ожидаем формат "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, `{"error": "invalid authorization format"}`, http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]
			claims := &Claims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, `{"error": "invalid or expired token"}`, http.StatusUnauthorized)
				return
			}

			// Проверяем срок действия вручную (дополнительный safeguard)
			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
				http.Error(w, `{"error": "token has expired"}`, http.StatusUnauthorized)
				return
			}

			// Сохраняем ID пользователя в контекст
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserIDFromContext безопасно извлекает ID пользователя из контекста в Usecase слое
func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(UserIDKey).(int)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return userID, nil
}
