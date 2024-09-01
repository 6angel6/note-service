package repository

import (
	"Zametki-go/internal/model"
	"database/sql"
)

type Authorization interface {
	Login(username, password string) (model.User, error)
}

type NoteRepository interface {
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
