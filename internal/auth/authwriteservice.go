package auth

import "fmt"

type AuthWriteService struct {
	repo AuthRepository
}

func NewWriteService(repo AuthRepository) *AuthWriteService {
	return &AuthWriteService{repo: repo}
}

func (s *AuthReadService) RegisterUser(username string, password string, rePassword string, email string) (bool, error) {
	if password != rePassword {

		return false, fmt.Errorf("password doesnt match")
	}
	hahsedPassword, err := HashPassword(password)
	if err != nil {

		return false, fmt.Errorf("something went wrong with hashing password : %s ", err)
	}
	result, err := s.repo.CreateUser(username, email, hahsedPassword)
	if err != nil {
		return false, fmt.Errorf("something went wrong with user creation : %s ", err)
	}
	return result, nil
}
