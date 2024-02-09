package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"start/internal/model"
	"start/internal/repository"
	"time"
)

const (
	salt       = "123u12p38yhfeduhfsdljbl71y23grb12pieb3asuhf7231w"
	signingKey = "sdgdfh#2342tssdtwe"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.AuthorizationRepository
}

func NewAuthService(repo repository.AuthorizationRepository) *AuthService {
	return &AuthService{repo: repo}
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

func (s AuthService) RegistrationService(user model.UserAuth) (string, error) {
	user.Password = generateHashPassword(user.Password)
	if len(user.Email) > 0 {
		return "", ErrorEmptyEmail
	}
	exist, err := s.repo.CheckUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		return "", ErrorCheckUser
	}
	if exist {
		return "", ErrorUsernameAlreadyExists
	}

	exist, err = s.repo.CheckUserByUsername(user.Username)
	if err != nil {
		return "", ErrorCheckUser
	}
	if exist {
		return "", ErrorDuplicateUsername
	}

	exist, err = s.repo.CheckUserByEmail(user.Email)
	if err != nil {
		return "", ErrorCheckUser
	}
	if exist {
		return "", ErrorDuplicateEmail
	}

	err = s.repo.CreateUserStorageAuth(user)
	if err != nil {
		return "", ErrorRegistrationUser
	}
	token, err := generateToken(user)
	if err != nil {
		return "", ErrorCreatedToken
	}

	return token, nil
}

func (s AuthService) LoginService(username, password string) (string, error) {
	exist, err := s.repo.CheckUserByUsernameAndPassword(username, generateHashPassword(password))
	if err != nil {
		return "", ErrorCheckUser
	}
	if !exist {
		return "", ErrorNotFoundUser
	}
	user := model.UserAuth{Username: username}
	token, err := generateToken(user)
	if err != nil {
		return "", ErrorCreatedToken
	}
	return token, nil
}
