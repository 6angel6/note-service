package repository

import (
	"Zametki-go/internal/model"
	"database/sql"
)

type NoteRepo struct {
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) *NoteRepo {
	return &NoteRepo{db: db}
}

func (r *NoteRepo) Create(note model.Note, userId string) (model.Note, error) {
	_, err := r.db.Exec("INSERT INTO notes (content, user_id) VALUES ($1, $2)", note.Content, userId)
	if err != nil {
		return model.Note{}, err
	}
	return note, nil
}

func (r *NoteRepo) GetAllNotesByUser(userId string) ([]model.Note, error) {
	rows, err := r.db.Query("SELECT note_id, content, user_id, created_at FROM notes WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.Id, &note.Content, &note.UserId, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
