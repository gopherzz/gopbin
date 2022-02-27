package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gopherzz/gopbin/internal/models"
)

const (
	defaultExtention  = ".txt"
	defaultNameFormat = "/%s"
)

type DocumentsRepo struct {
	docsFolderPath string
	pathFormat     string
}

func NewDocumentsRepo(docsFolderPath string) *DocumentsRepo {
	return &DocumentsRepo{
		docsFolderPath: docsFolderPath,
		pathFormat:     docsFolderPath + defaultNameFormat + defaultExtention,
	}
}

func (r *DocumentsRepo) formatDocPath(id string) string {
	return fmt.Sprintf(r.pathFormat, id)
}

func (r *DocumentsRepo) Create(doc *models.Document) (err error) {
	f, err := os.Create(r.formatDocPath(doc.ID))
	if err != nil {
		return
	}

	_, err = f.WriteString(doc.Content)
	if err != nil {
		return
	}

	return nil
}

func (r *DocumentsRepo) Read(id string) (*models.Document, error) {
	f, err := os.Open(r.formatDocPath(id))
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &models.Document{
		ID:      id,
		Content: string(bytes),
	}, nil
}
