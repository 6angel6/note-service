package model

type Note struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
