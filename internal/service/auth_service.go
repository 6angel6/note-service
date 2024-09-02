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
		return Tokens{}, errors.New("unauthorized")
	}

	token, err := jwt.CreateToken(user.Username)
	if err != nil {
		log.Println("Error creating access token:", err)
		return Tokens{}, err
	}

	refreshToken, err := jwt.CreateRefreshToken(user.Username)
	if err != nil {
		log.Println("Error creating refresh token:", err)
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (s *AuthService) Refresh(refreshToken string) (Tokens, error) {
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		log.Println("Error validating refresh token:", err)
		return Tokens{}, errors.New("unauthorized")
	}
	username := claims["username"].(string)

	accessToken, err := jwt.CreateToken(username)
	if err != nil {
		log.Println("Error creating new access token:", err)
		return Tokens{}, err
	}

	newRefreshToken, err := jwt.CreateRefreshToken(username)
	if err != nil {
		log.Println("Error creating new refresh token:", err)
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}

	return tokens, nil
}

func (s *AuthService) GetUserIdByUsername(username string) (string, error) {
	return s.repo.GetUserIDByUsername(username)
}
