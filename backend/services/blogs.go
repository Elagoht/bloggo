package services

import (
	"net/http"

	"github.com/Elagoht/bloggo/backend/models"
	"github.com/Elagoht/bloggo/backend/pipes"
	"github.com/Elagoht/bloggo/backend/repositories"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-playground/validator/v10"
)

type BlogService struct {
	repository *repositories.BlogRepository
	validate   *validator.Validate
}

func NewBlogService() *BlogService {
	return &BlogService{
		repository: repositories.NewBlogRepository(),
		validate:   pipes.GetValidator(),
	}
}

func (service *BlogService) GetAll() ([]models.ResponseBlogCard, *utils.AppError) {
	return service.repository.GetAll()
}

func (service *BlogService) GetBySlug(slug string) (models.ResponseBlog, *utils.AppError) {
	return service.repository.GetBySlug(slug)
}

func (service *BlogService) CreateBlog(blog models.RequestBlog) *utils.AppError {
	err := service.validate.Struct(blog)
	if err != nil {
		return utils.NewAppError(
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
	}
	return service.repository.CreateBlog(blog)
}

func (service *BlogService) PublishBlog(slug string) *utils.AppError {
	return service.repository.PublishBlog(slug)
}

func (service *BlogService) UnpublishBlog(slug string) *utils.AppError {
	return service.repository.UnpublishBlog(slug)
}

func (service *BlogService) UpdateBlog(slug string, blog models.RequestBlogPartial) *utils.AppError {
	return service.repository.PatchBlog(slug, blog)
}

func (service *BlogService) DeleteBlog(slug string) *utils.AppError {
	return service.repository.DeleteBlog(slug)
}
