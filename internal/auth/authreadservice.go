package auth

import (
	"fmt"
	"virtual-diary/internal/auth/userdao"
	"virtual-diary/internal/auth/userdto"
)

type ReadService struct {
	repo AuthRepository
}

func NewReadService(repo AuthRepository) *ReadService {
	return &ReadService{repo: repo}
}

func convertDaoToDto(userDTO *userdto.UserDTO, userDAO userdao.User) {
	userDTO.Email = userDAO.Email
	userDTO.Username = userDAO.Username
}

func (s *ReadService) GetUserByUsername(username string) (userdto.UserDTO, error) {
	var userDTO userdto.UserDTO
	userDAO, err := s.repo.GetUserByUsername(username)
	if err != nil {
		fmt.Errorf("Error will be here")
		return userdto.UserDTO{}, err

	}
	convertDaoToDto(&userDTO, userDAO)

	return userDTO, nil
}

func (s *ReadService) SignIn(username string, password string) (string, error) {
	userDAO, err := s.repo.GetUserByUsername(username)
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
