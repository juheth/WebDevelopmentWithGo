package services

import (
	"github.com/juheth/WebDevelopmentWithGo.git/models"
	"github.com/juheth/WebDevelopmentWithGo.git/repository"
)

func SavePage(p *models.Page) error {
	return repository.SavePageToFile(p)
}

func LoadPage(title string) (*models.Page, error) {
	return repository.LoadPage(title)
}
