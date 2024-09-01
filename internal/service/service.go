package service

import "Zametki-go/internal/repository"

type Authorization interface {
	Login(username, password string) (Tokens, error)
	Refresh(refreshToken string) (Tokens, error)
}

type Note interface {
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
