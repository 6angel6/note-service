package service

import (
	"Zametki-go/internal/model"
	"Zametki-go/internal/model/dto"
	"Zametki-go/internal/repository"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(note dto.NoteRequest, userId string) error {
	n := model.Note{
		UserId:  userId,
		Content: note.Content,
	}
	_, err := s.repo.Create(n, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) GetAllNotes(userId string) ([]model.Note, error) {
	notes, err := s.repo.GetAllNotesByUser(userId)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
