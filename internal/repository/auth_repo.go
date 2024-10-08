package repository

import (
	"Zametki-go/internal/model"
	"database/sql"
	"errors"
	"fmt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) Login(username, password string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT username, password FROM users WHERE username = $1 AND password = $2", username, password).Scan(&user.Username, &user.Password)
	if err != nil {
		return model.User{}, errors.New("invalid username or password")
	}
	return user, nil
}

func (r *AuthRepo) GetUserIDByUsername(username string) (string, error) {
	var userId string
	err := r.db.QueryRow("SELECT user_id FROM users WHERE username = $1", username).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	return userId, nil
}
