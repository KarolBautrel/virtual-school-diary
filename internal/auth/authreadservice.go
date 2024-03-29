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
		fmt.Errorf("error will be here: %s", err)
		return userdto.UserDTO{}, err

	}
	ConvertUserDaoToDto(&userDTO, userDAO)

	return userDTO, nil
}

func (s *AuthReadService) SignIn(username string, password string) (string, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	userDAO, err := s.repo.GetUserByUsername(username, ctx)
	if err != nil {
		fmt.Errorf("error will be here: %s", err)
		return "", err
	}
	if !VerifyPassword(userDAO.Password, password) {
		fmt.Errorf("error will be here: %s", err)
		return "", err
	}
	token, err := CreateToken(userDAO.ID)
	if err != nil {
		fmt.Errorf("error will be here: %s", err)
		return "", err
	}
	return token, nil
}
