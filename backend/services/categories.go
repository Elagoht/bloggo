package services

import (
	"github.com/Elagoht/bloggo/backend/models"
	"github.com/Elagoht/bloggo/backend/pipes"
	"github.com/Elagoht/bloggo/backend/repositories"
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
) (models.ResponseCategory, *utils.AppError) {
	category, err := service.repository.GetBySlug(slug)
	if err != nil {
		return models.ResponseCategory{}, err
	}

	return category, nil
}

func (service *CategoryService) Create(
	category models.RequestCategory,
) *utils.AppError {
	validationErr := service.validate.Struct(category)
	if validationErr != nil {
		return utils.ErrBadRequest
	}

	err := service.repository.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (service *CategoryService) Patch(
	slug string,
	category models.RequestCategoryPartial,
) *utils.AppError {
	validationErr := service.validate.Struct(category)
	if validationErr != nil {
		return utils.ErrBadRequest
	}

	return service.repository.PatchCategory(slug, category)
}

func (service *CategoryService) Delete(slug string) *utils.AppError {
	return service.repository.SoftDeleteCategory(slug)
}
