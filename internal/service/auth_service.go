package service

import (
	"Zametki-go/internal/repository"
	"Zametki-go/pkg/jwt"
	"errors"
	"log"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	repo   repository.Authorization
	Tokens Tokens
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(username, password string) (Tokens, error) {
	user, err := s.repo.Login(username, password)
	if err != nil {
		log.Println("Error authenticate user:", err)
		return Tokens{}, errors.New("unauthorized") // Возвращаем ошибку при неудачной аутентификации
	}

	token, err := jwt.CreateToken(user.Username)
	if err != nil {
		log.Println("Error creating access token:", err)
		return Tokens{}, err // Возвращаем ошибку, если не удалось создать токен
	}

	refreshToken, err := jwt.CreateRefreshToken(user.Username)
	if err != nil {
		log.Println("Error creating refresh token:", err)
		return Tokens{}, err // Возвращаем ошибку, если не удалось создать рефреш-токен
	}

	tokens := Tokens{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}
