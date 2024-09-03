package dto

type NoteRequest struct {
	Content string `json:"content" validate:"required,min=5"`
}
