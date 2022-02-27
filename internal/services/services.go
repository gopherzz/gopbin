package services

import (
	"github.com/gopherzz/gopbin/internal/models"
	"github.com/gopherzz/gopbin/internal/repository"
)

type Documents interface {
	CreateDocument(docReq *models.DocumentRequest) (id string, err error)
	ReadDocument(id string) (*models.Document, error)
}

type Services struct {
	Documents
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Documents: NewDocumentsService(repo),
	}
}
