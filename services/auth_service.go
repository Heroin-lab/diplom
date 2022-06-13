package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	salt       = "9dlk1jfh5aj1er2ui4fh"
	signingKey = "putin huilo"
)

type AuthService struct {
	repo repositories.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repositories.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user request.User) (int, error) {
	newUser := repositories.User{
		Login:    user.Login,
		Password: user.Password,
	}

	newUser.Password = s.generatePasswordHash(user.Password)

	userId, err := s.repo.CreateUser(newUser)
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		logrus.Errorf("User with login '%s'  already exist!", newUser.Login)
		return 0, errors.New("User with this login already exist!")
	}
	return userId, err
}

func (s *AuthService) GenerateToken(login string, password string, duration int) (string, error) {
	user, err := s.repo.GetUserByLogin(login, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(duration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.User_id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims is invalid")
	}

	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
