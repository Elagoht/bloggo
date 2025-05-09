package repositories

import (
	"database/sql"

	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/utils"
)

type BlogRepository struct {
	dataBase *sql.DB
}

func NewBlogRepository() *BlogRepository {
	return &BlogRepository{
		dataBase: utils.DataBase,
	}
}

func (repository *BlogRepository) GetAll() ([]models.ResponseBlogCard, *utils.AppError) {
	rows, err := repository.dataBase.Query(SQLGetAllBlogs)
	if err != nil {
		return nil, utils.MapDatabaseError(err)
	}
	defer rows.Close()

	blogs := []models.ResponseBlogCard{}

	for rows.Next() {
		var blog models.ResponseBlogCard
		err = rows.Scan(
			&blog.Title, &blog.Slug, &blog.Spot, &blog.Cover, &blog.ReadTime,
			&blog.ReadCount, &blog.PublishedAt, &blog.CategoryName,
			&blog.CategorySlug,
		)
		if err != nil {
			return nil, utils.MapDatabaseError(err)
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repository *BlogRepository) GetBySlug(
	slug string,
) (models.ResponseBlog, *utils.AppError) {
	row := repository.dataBase.QueryRow(SQLGetBlogBySlug, slug)

	var blog models.ResponseBlog
	err := row.Scan(
		&blog.Slug, &blog.Title, &blog.Content, &blog.Keywords, &blog.Description,
		&blog.Spot, &blog.Cover, &blog.PublishedAt, &blog.CategoryName,
		&blog.CategorySlug,
	)
	if err != nil {
		return models.ResponseBlog{}, utils.MapDatabaseError(err)
	}

	return blog, nil
}

func (repository *BlogRepository) CreateBlog(
	blog models.RequestBlog,
) *utils.AppError {
	statement, err := repository.dataBase.Prepare(SQLCreateBlog)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		blog.Title, blog.Slug, blog.Spot, blog.Content, blog.Keywords, blog.Description,
		blog.Cover, blog.CalculateReadTime(), blog.Published, blog.CategoryId,
		blog.CalculatePublishedAt(),
	)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	return nil
}

func (repository *BlogRepository) changePublishedStatus(
	slug string,
	published bool,
) *utils.AppError {
	statement, err := repository.dataBase.Prepare(SQLChangePublishedStatus)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(published, slug)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	return nil
}

func (repository *BlogRepository) PublishBlog(
	slug string,
) *utils.AppError {
	return repository.changePublishedStatus(slug, true)
}

func (repository *BlogRepository) UnpublishBlog(
	slug string,
) *utils.AppError {
	return repository.changePublishedStatus(slug, false)
}

func (repository *BlogRepository) PatchBlog(
	slug string,
	blog models.RequestBlogPartial,
) *utils.AppError {

	fields := []utils.PatchField{
		{Name: "title", Value: blog.Title, Skip: blog.Title == ""},
		{Name: "content", Value: blog.Content, Skip: blog.Content == ""},
		{Name: "keywords", Value: blog.Keywords, Skip: blog.Keywords == ""},
		{Name: "description", Value: blog.Description, Skip: blog.Description == ""},
		{Name: "spot", Value: blog.Spot, Skip: blog.Spot == ""},
		{Name: "cover", Value: blog.Cover, Skip: blog.Cover == ""},
		{Name: "published", Value: blog.Published, Skip: !blog.Published},
		{Name: "categoryId", Value: blog.CategoryId, Skip: blog.CategoryId == 0},
		{Name: "updatedAt", Value: "datetime('now')", Skip: false},
	}

	query, args := utils.BuildPatchQuery("blogs", "slug", slug, fields)

	statement, err := repository.dataBase.Prepare(query)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(args...)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	return nil
}

func (repository *BlogRepository) DeleteBlog(
	slug string,
) *utils.AppError {
	statement, err := repository.dataBase.Prepare(SQLSoftDeleteBlog)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(slug)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	return nil
}
