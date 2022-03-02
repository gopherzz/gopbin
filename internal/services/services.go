package services

import (
	"github.com/gopherzz/gopbin/internal/models"
	"github.com/gopherzz/gopbin/internal/repository"
	"github.com/gopherzz/gopbin/pkg/id"
)

type Documents interface {
	CreateDocument(docReq *models.DocumentRequest) (id id.ID, err error)
	ReadDocument(id id.ID) (*models.Document, error)
}

type Services struct {
	Documents
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Documents: NewDocumentsService(repo),
	}
}
