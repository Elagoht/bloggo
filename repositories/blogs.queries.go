package repositories

const (
	// Get all non-deleted blogs
	SQLGetAllBlogs = `
		SELECT b.title, b.slug, b.spot, b.cover, b.readTime, b.readCount, b.publishedAt, c.name, c.slug
		FROM blogs b
		INNER JOIN categories c ON b.categoryId = c.id
		WHERE b.deletedAt IS NULL AND c.deletedAt IS NULL
		ORDER BY b.createdAt DESC
	`
	// Get a blog by slug if it is not deleted
	SQLGetBlogBySlug = `
		SELECT b.title, b.slug, b.spot, b.cover, b.readTime, b.readCount, b.publishedAt, c.name, c.slug
		FROM blogs b
		INNER JOIN categories c ON b.categoryId = c.id
		WHERE b.slug = ? AND b.deletedAt IS NULL AND c.deletedAt IS NULL
	`
	// Create a new blog, requires manually setting readTime, publishedAt and categoryId
	SQLCreateBlog = `
		INSERT INTO blogs (title, slug, spot, content, keywords, description, cover, readTime, published, categoryId, publishedAt, updatedAt)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, datetime('now'))
	`
	// Publish a blog
	SQLPublishBlog = `
		UPDATE blogs
		SET published = 1, publishedAt = datetime('now'), updatedAt = datetime('now')
		WHERE slug = ?
	`
	// Unpublish a blog
	SQLUnpublishBlog = `
		UPDATE blogs
		SET published = 0, updatedAt = datetime('now')
		WHERE slug = ?
	`
	// Soft delete a blog
	SQLSoftDeleteBlog = `
		UPDATE blogs
		SET deletedAt = datetime('now'), updatedAt = datetime('now')
		WHERE slug = ?
	`
)
