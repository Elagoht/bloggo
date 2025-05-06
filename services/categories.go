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
