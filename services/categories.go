package services

import (
	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/repositories"
)

type CategoryService struct {
	repository *repositories.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repository: repositories.NewCategoryRepository(),
	}
}

func (service *CategoryService) GetAll() ([]models.Category, error) {
	return service.repository.GetAll()
}

func (service *CategoryService) GetBySlug(slug string) (models.Category, error) {
	category, err := service.repository.GetBySlug(slug)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}
