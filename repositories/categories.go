package repositories

import (
	"database/sql"

	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/utils"
	"github.com/mattn/go-sqlite3"
)

type CategoryRepository struct {
	dataBase *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		dataBase: utils.DataBase,
	}
}

func (repository *CategoryRepository) GetAll() ([]models.ResponseCategoryListItem, error) {
	rows, err := repository.dataBase.Query("SELECT id, name, slug FROM categories ORDER BY createdAt DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []models.ResponseCategoryListItem{}

	for rows.Next() {
		var category models.ResponseCategoryListItem
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Slug,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepository) GetBySlug(slug string) (models.Category, error) {
	row := repository.dataBase.QueryRow("SELECT * FROM categories WHERE slug = ?", slug)

	var category models.Category

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.Keywords,
		&category.Spot,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		return models.Category{}, utils.NewError("category not found")
	}

	return category, nil
}

func (repository *CategoryRepository) Create(
	category models.RequestCategoryCreate,
) (models.RequestCategoryCreate, error) {
	statement, err := repository.dataBase.Prepare(
		"INSERT INTO categories (name, slug, description, keywords, spot) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return models.RequestCategoryCreate{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		category.Name,
		category.Slug,
		category.Description,
		category.Keywords,
		category.Spot,
	)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return models.RequestCategoryCreate{}, utils.NewError("category already exists")
		}
		return models.RequestCategoryCreate{}, utils.NewError("failed to create category")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return models.RequestCategoryCreate{}, err
	}

	category.Id = int(lastInsertId)

	return category, nil
}

func (repository *CategoryRepository) Update(
	slug string,
	category models.RequestCategoryUpdate,
) (models.RequestCategoryUpdate, error) {
	statement, err := repository.dataBase.Prepare(
		"UPDATE categories SET name = ?, slug = ?, description = ?, keywords = ?, spot = ? WHERE slug = ?",
	)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(
		category.Name,
		category.Slug,
		category.Description,
		category.Keywords,
		category.Spot,
		slug,
	)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	query, args := utils.UpdateUpdatedAtQuery("categories", "slug", slug)
	_, err = repository.dataBase.Exec(query, args...)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	return category, nil
}

func (repository *CategoryRepository) Patch(
	slug string,
	category models.RequestCategoryUpdate,
) (models.RequestCategoryUpdate, error) {
	fields := []utils.PatchField{
		{Name: "name", Value: category.Name, Skip: category.Name == ""},
		{Name: "slug", Value: category.Slug, Skip: category.Slug == ""},
		{Name: "description", Value: category.Description, Skip: category.Description == ""},
		{Name: "keywords", Value: category.Keywords, Skip: category.Keywords == ""},
		{Name: "spot", Value: category.Spot, Skip: category.Spot == ""},
	}

	query, args := utils.BuildPatchQuery("categories", "slug", slug, fields)

	statement, err := repository.dataBase.Prepare(query)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(args...)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	if rowsAffected == 0 {
		return models.RequestCategoryUpdate{}, utils.NewError("category not found")
	}

	updatedCategory, err := repository.GetBySlug(slug)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	query, args = utils.UpdateUpdatedAtQuery("categories", "slug", slug)
	_, err = repository.dataBase.Exec(query, args...)
	if err != nil {
		return models.RequestCategoryUpdate{}, err
	}

	return models.RequestCategoryUpdate{Category: updatedCategory}, nil
}

func (repository *CategoryRepository) Delete(
	slug string,
) error {
	result, err := repository.dataBase.Exec(
		"DELETE FROM categories WHERE slug = ?",
		slug,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.NewError("category not found")
	}

	return nil
}
