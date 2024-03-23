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
		fmt.Errorf("Error will be here")
		return userdto.UserDTO{}, err

	}
	ConvertDaoToDto(&userDTO, userDAO)

	return userDTO, nil
}

func (s *AuthReadService) SignIn(username string, password string) (string, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	userDAO, err := s.repo.GetUserByUsername(username, ctx)
	if err != nil {
		fmt.Errorf("Error will be here")
		return "", err
	}
	if !VerifyPassword(userDAO.Password, password) {
		fmt.Errorf("Error will be here")
		return "", err
	}
	token, err := CreateToken(userDAO.ID)
	if err != nil {
		fmt.Errorf("Error will be here")
		return "", err
	}
	return token, nil
}
