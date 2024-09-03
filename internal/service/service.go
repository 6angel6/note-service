package service

import (
	"Zametki-go/internal/model"
	"Zametki-go/internal/model/dto/request"
	"Zametki-go/internal/repository"
)

type Authorization interface {
	Login(username, password string) (Tokens, error)
	Refresh(refreshToken string) (Tokens, error)
	GetUserIdByUsername(username string) (string, error)
}

type Note interface {
	Create(note dto.NoteRequest, userId string) error
	GetAllNotes(userId string) ([]model.Note, error)
}

type Service struct {
	Authorization
	Note
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Note:          NewNoteService(r.NoteRepository),
	}
}
