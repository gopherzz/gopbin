package repository

import (
	"github.com/gopherzz/gopbin/internal/models"
	"github.com/gopherzz/gopbin/pkg/id"
)

type Documents interface {
	Create(doc *models.Document) (err error)
	Read(id id.ID) (*models.Document, error)
}

type Repository struct {
	Documents
}

func NewRepository(documentsFolderPath string) *Repository {
	return &Repository{
		Documents: NewDocumentsRepo(documentsFolderPath),
	}
}
