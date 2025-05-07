package models

import "time"

type Blog struct {
	Id int `json:"id"`

	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`

	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Spot        string `json:"spot"`
	Cover       string `json:"cover"`

	Published bool `json:"published"`

	CreatedAt time.Time `json:"createdAt" default:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" default:"CURRENT_TIMESTAMP"`
	ReleaseAt time.Time `json:"releaseAt"`
	DeletedAt time.Time `json:"deletedAt"`

	ReadTime  int `json:"readTime"`
	ReadCount int `json:"readCount"`

	CategoryId int `json:"categoryId"`
}

type BlogWithCategory struct {
	Blog
	Category Category `json:"category"`
}

type RequestBlogCreate struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`

	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Spot        string `json:"spot"`
	Cover       string `json:"cover"`

	Published bool `json:"published"`

	CategoryId int `json:"categoryId"`
}

type RequestBlogUpdate struct {
	RequestBlogCreate
}
