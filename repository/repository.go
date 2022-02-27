package repository

import "github.com/gopherzz/gopbin/internal/models"

type Documents interface {
	Create(doc *models.Document) (err error)
	Read(id string) (*models.Document, error)
}

type Repository struct {
	Documents
}

func NewRepository(documentsFolderPath string) *Repository {
	return &Repository{
		Documents: NewDocumentsRepo(documentsFolderPath),
	}
}
