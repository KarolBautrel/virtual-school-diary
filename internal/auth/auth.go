package auth

import "virtual-diary/internal/auth/userdao"

type AuthWriteService interface {
}

type AuthReadService interface {
}

type AuthRepository interface {
	CreateUser(username string, email string, password string, rePassword string) (bool, error)
	GetUserByUsername(username string) (userdao.User, error)
}
