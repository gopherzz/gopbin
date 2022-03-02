package models

import "github.com/gopherzz/gopbin/pkg/id"

type DocumentRequest struct {
	Content string `json:"content"`
}

type Document struct {
	ID      id.ID  `json:"id"`
	Content string `json:"content"`
}
