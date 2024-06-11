package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" validate:"required,max=100"`
	Author string `json:"author" validate:"required,max=10"`
}
