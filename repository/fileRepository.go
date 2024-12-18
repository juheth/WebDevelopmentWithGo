package repository

import (
	"log"
	"os"
	"path/filepath"

	"github.com/juheth/WebDevelopmentWithGo.git/models"
)

const dataDir = "data"

func init() {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatal(err)
	}
}

func SavePageToFile(p *models.Page) error {
	filename := filepath.Join(dataDir, p.Title+".txt")
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*models.Page, error) {
	filename := filepath.Join(dataDir, title+".txt")
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}
