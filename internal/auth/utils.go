package auth

import (
	"time"
	"virtual-diary/internal/auth/userdao"
	"virtual-diary/internal/auth/userdto"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CreateToken(userId uint) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("WillBeSomeSecretKey"))

	if err != nil {
		return "", err
	}
	return token, nil
}

func ConvertDaoToDto(userDTO *userdto.UserDTO, userDAO userdao.User) {
	userDTO.Email = userDAO.Email
	userDTO.Username = userDAO.Username
}
