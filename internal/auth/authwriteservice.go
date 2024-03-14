package auth

import "fmt"

type AuthWriteService struct {
	repo AuthRepository
}

func NewWriteService(repo AuthRepository) *AuthWriteService {
	return &AuthWriteService{repo: repo}
}

func (s *AuthReadService) RegisterUser(username string, password string, rePassword string, email string) bool {
	if password != rePassword {
		fmt.Errorf("Error will be here")
		return false
	}
	hahsedPassword, err := HashPassword(password)
	if err != nil {
		fmt.Errorf("Error will be here")
		return false
	}
	result, err := s.repo.CreateUser(username, email, hahsedPassword)
	if err != nil {
		return false
	}
	return result
}
