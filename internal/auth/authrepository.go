package auth

import (
	"virtual-diary/internal/auth/userdao"

	"gorm.io/gorm"
)

type authRepoImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepoImpl{db: db}
}

func (a *authRepoImpl) CreateUser(username string, email string, password string, rePassword string) (bool, error) {
	newUser := userdao.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	result := a.db.Create(&newUser)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (a *authRepoImpl) GetUserByUsername(username string) (userdao.User, error) {
	var user userdao.User
	result := a.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return userdao.User{}, result.Error
	}
	return user, nil
}
