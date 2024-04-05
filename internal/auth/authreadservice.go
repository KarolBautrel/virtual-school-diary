package auth

import (
	"fmt"
	"time"
	"virtual-diary/internal/auth/userdto"
	globalutils "virtual-diary/pkg/utils"
)

type AuthReadService struct {
	repo AuthRepository
}

func NewReadService(repo AuthRepository) *AuthReadService {
	return &AuthReadService{repo: repo}
}

func (s *AuthReadService) GetUserByUsername(username string) (userdto.UserDTO, error) {
	var userDTO userdto.UserDTO
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	userDAO, err := s.repo.GetUserByUsername(username, ctx)
	if err != nil {

		return userdto.UserDTO{}, fmt.Errorf("error with getting user: %s", err)

	}
	ConvertUserDaoToDto(&userDTO, userDAO)

	return userDTO, nil
}

func (s *AuthReadService) SignIn(username string, password string) (string, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	userDAO, err := s.repo.GetUserByUsername(username, ctx)
	if err != nil {
		return "", fmt.Errorf("error with getting username from repo: %s", err)

	}
	if !VerifyPassword(userDAO.Password, password) {

		return "", fmt.Errorf("password does not match: %s", err)
	}
	token, err := CreateToken(userDAO.ID)
	if err != nil {

		return "", fmt.Errorf("there is an error with token: %s", err)
	}
	return token, nil
}
