package repository

import (
	"log"
	"os"

	"github.com/juheth/WebDevelopmentWithGo.git/models"
)

func SavePageToFile(p *models.Page) error {
	fileName := p.Title + ".txt"
	err := os.WriteFile(fileName, p.Body, 0600)
	if err != nil {
		log.Printf("Failed to write file %s: %v", fileName, err)
		return err
	}
	return nil
}

func LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}
