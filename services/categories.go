package services

import (
	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/pipes"
	"github.com/Elagoht/bloggo/repositories"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-playground/validator/v10"
)

type CategoryService struct {
	repository *repositories.CategoryRepository
	validate   *validator.Validate
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repository: repositories.NewCategoryRepository(),
		validate:   pipes.GetValidator(),
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
	category models.RequestCategory,
) (models.RequestCategory, *utils.AppError) {
	validationErr := service.validate.Struct(category)
	if validationErr != nil {
		return models.RequestCategory{}, utils.ErrBadRequest
	}

	createdCategory, err := service.repository.Create(category)
	if err != nil {
		return models.RequestCategory{}, err
	}

	return createdCategory, nil
}

func (service *CategoryService) Update(
	slug string,
	category models.RequestCategory,
) *utils.AppError {
	validationErr := service.validate.Struct(category)
	if validationErr != nil {
		return utils.ErrBadRequest
	}

	return service.repository.Update(slug, category)
}

func (service *CategoryService) Patch(
	slug string,
	category models.RequestCategoryPartial,
) *utils.AppError {
	validationErr := service.validate.Struct(category)
	if validationErr != nil {
		return utils.ErrBadRequest
	}

	return service.repository.Patch(slug, category)
}

func (service *CategoryService) Delete(slug string) *utils.AppError {
	return service.repository.Delete(slug)
}
