package services

import (
	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/repositories"
	"github.com/Elagoht/bloggo/utils"
)

type CategoryService struct {
	repository *repositories.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repository: repositories.NewCategoryRepository(),
	}
}

func (service *CategoryService) GetAll() ([]models.ResponseCategoryListItem, *utils.AppError) {
	return service.repository.GetAll()
}

func (service *CategoryService) GetBySlug(
	slug string,
) (models.Category, *utils.AppError) {
	category, err := service.repository.GetBySlug(slug)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (service *CategoryService) Create(
	category models.RequestCategoryCreate,
) (models.RequestCategoryCreate, *utils.AppError) {
	createdCategory, err := service.repository.Create(category)
	if err != nil {
		return models.RequestCategoryCreate{}, err
	}

	return createdCategory, nil
}

func (service *CategoryService) Update(
	slug string,
	category models.RequestCategoryUpdate,
) (models.RequestCategoryUpdate, *utils.AppError) {
	return service.repository.Update(slug, category)
}

func (service *CategoryService) Patch(
	slug string,
	category models.RequestCategoryUpdate,
) (models.RequestCategoryUpdate, *utils.AppError) {
	return service.repository.Patch(slug, category)
}

func (service *CategoryService) Delete(slug string) *utils.AppError {
	return service.repository.Delete(slug)
}
