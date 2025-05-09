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
	rows, err := repository.dataBase.Query(`
		SELECT b.title, b.slug, b.spot, b.cover, b.readTime, b.readCount, b.publishedAt, c.name, c.slug
		FROM blogs b
		INNER JOIN categories c ON b.categoryId = c.id
		WHERE b.deletedAt IS NULL
		ORDER BY b.createdAt DESC
	`)
	if err != nil {
		return nil, utils.MapDatabaseError(err)
	}
	defer rows.Close()

	blogs := []models.ResponseBlogCard{}

	for rows.Next() {
		var blog models.ResponseBlogCard
		err = rows.Scan(
			&blog.Title,
			&blog.Slug,
			&blog.Spot,
			&blog.Cover,
			&blog.ReadTime,
			&blog.ReadCount,
			&blog.PublishedAt,
			&blog.CategoryName,
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
	row := repository.dataBase.QueryRow(`
		SELECT b.title, b.slug, b.spot, b.cover, b.readTime, b.readCount, b.publishedAt, c.name, c.slug
		FROM blogs b
		INNER JOIN categories c ON b.categoryId = c.id
		WHERE b.slug = ?
	`, slug)

	var blog models.ResponseBlog
	err := row.Scan(
		&blog.Slug,
		&blog.Title,
		&blog.Content,
		&blog.Keywords,
		&blog.Description,
		&blog.Spot,
		&blog.Cover,
		&blog.PublishedAt,
		&blog.CategoryName,
		&blog.CategorySlug,
	)
	if err != nil {
		return models.ResponseBlog{}, utils.MapDatabaseError(err)
	}

	return blog, nil
}
