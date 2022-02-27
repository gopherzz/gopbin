package services

import (
	"github.com/gopherzz/gopbin/internal/models"
	"github.com/gopherzz/gopbin/pkg/id"
	"github.com/gopherzz/gopbin/internal/repository"
)

type DocumentsService struct {
	repo *repository.Repository
}

func NewDocumentsService(repo *repository.Repository) *DocumentsService {
	return &DocumentsService{
		repo: repo,
	}
}

func (s *DocumentsService) CreateDocument(docReq *models.DocumentRequest) (docId string, err error) {
	docId = id.GenerateId()

	return docId, s.repo.Create(&models.Document{
		ID:   docId,
		Content: docReq.Content,
	})
}

func (s *DocumentsService) ReadDocument(id string) (*models.Document, error) {
	return s.repo.Read(id)
}
