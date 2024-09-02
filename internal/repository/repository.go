package repository

import (
	"Zametki-go/internal/model"
	"database/sql"
)

type Authorization interface {
	Login(username, password string) (model.User, error)
	GetUserIDByUsername(username string) (string, error)
}

type NoteRepository interface {
	Create(note model.Note, userId string) (model.Note, error)
	GetAllNotesByUser(userId string) ([]model.Note, error)
}

type Repository struct {
	Authorization
	NoteRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthRepo(db),
		NoteRepository: NewNoteRepo(db),
	}
}
