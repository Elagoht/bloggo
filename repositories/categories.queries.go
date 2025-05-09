package repositories

const (
	// Get all non-deleted categories
	SQLGetAllCategories = `
		SELECT name, slug FROM categories
		WHERE deletedAt IS NULL
		ORDER BY createdAt DESC
	`
	// Get a category by slug if it is not deleted
	SQLGetCategoryBySlug = `
		SELECT name, slug, description, keywords, spot
		FROM categories
		WHERE deletedAt IS NULL AND slug = ?
	`
	// Create a new category
	SQLCreateCategory = `
		INSERT INTO categories (name, slug, description, keywords, spot)
		VALUES (?, ?, ?, ?, ?)
	`
	// Update a category, by taking whole object, patch request will use a query generator
	SQLUpdateCategory = `
		UPDATE categories
		SET name = ?, slug = ?, description = ?, keywords = ?, spot = ?, updatedAt = datetime('now')
		WHERE slug = ?
	`
	// Soft delete a category
	SQLSoftDeleteCategory = `
		UPDATE categories
		SET deletedAt = datetime('now'), updatedAt = datetime('now')
		WHERE slug = ?
	`
)
