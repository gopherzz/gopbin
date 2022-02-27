package models

type DocumentRequest struct {
	Content string `json:"content"`
}

type Document struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
