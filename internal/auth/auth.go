package auth

import (
	"context"
	"virtual-diary/internal/auth/userdao"
)

type AuthRepository interface {
	CreateUser(username string, email string, password string) (bool, error)
	GetUserByUsername(username string, ctx context.Context) (userdao.User, error)
	GetUserById(userId string, ctx context.Context) (userdao.User, error)
}
