package services

import (
	"os"

	"github.com/juheth/WebDevelopmentWithGo.git/models"
)

func SavePage(p *models.Page) error {
	filename := p.Title + ".txt"
	return os.WriteFile(
		filename,
		p.Body,
		0600,
	)
}

func LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}
