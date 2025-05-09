package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/services"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-chi/chi"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: services.NewCategoryService(),
	}
}

func (controller *CategoryController) GetAllCategories(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	categories, err := controller.categoryService.GetAll()
	if err != nil {
		return err
	}

	json.NewEncoder(writer).Encode(categories)
	return nil
}

func (controller *CategoryController) GetCategoryBySlug(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	category, err := controller.categoryService.GetBySlug(slug)
	if err != nil {
		return err
	}

	json.NewEncoder(writer).Encode(category)
	return nil
}

func (controller *CategoryController) CreateCategory(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	var category models.RequestCategory
	bodyErr := json.NewDecoder(request.Body).Decode(&category)
	if bodyErr != nil {
		return utils.ErrInternalServer
	}

	createdCategory, err := controller.categoryService.Create(category)
	if err != nil {
		return err
	}

	json.NewEncoder(writer).Encode(createdCategory)
	return nil
}

func (controller *CategoryController) UpdateCategory(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	var category models.RequestCategory
	bodyErr := json.NewDecoder(request.Body).Decode(&category)
	if bodyErr != nil {
		return utils.ErrInternalServer
	}

	err := controller.categoryService.Update(slug, category)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusNoContent)
	return nil
}

func (controller *CategoryController) PatchCategory(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	var category models.RequestCategoryPartial
	bodyErr := json.NewDecoder(request.Body).Decode(&category)
	if bodyErr != nil {
		return utils.ErrInternalServer
	}

	err := controller.categoryService.Patch(slug, category)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusNoContent)
	return nil
}

func (controller *CategoryController) DeleteCategory(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	err := controller.categoryService.Delete(slug)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusNoContent)
	return nil
}
