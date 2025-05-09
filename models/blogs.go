package models

import (
	"strings"
	"time"
)

const (
	wordsPerMinute = 180
)

type Blog struct {
	// Identifiers
	Id   int    `json:"id" validate:"required"`
	Slug string `json:"slug" validate:"required,max=100,slug"`
	// Content and metadata
	Title       string `json:"title" validate:"required,max=100"`
	Content     string `json:"content" validate:"required"`
	Keywords    string `json:"keywords" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=155"`
	Spot        string `json:"spot" validate:"required,max=75"`
	Cover       string `json:"cover" validate:"required,url"`
	// Status
	Published bool `json:"published" validate:"required"`
	// Timestamps
	CreatedAt   time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt" validate:"required"`
	PublishedAt time.Time  `json:"publishedAt" validate:"required"`
	DeletedAt   *time.Time `json:"deletedAt"`
	// Statistics
	ReadTime  int `json:"readTime" validate:"required"`
	ReadCount int `json:"readCount" validate:"required"`
	// Category ID
	CategoryId int `json:"categoryId" validate:"required"`
}

type RequestBlog struct {
	Title       string `json:"title" validate:"required,max=100"`
	Slug        string `json:"slug" validate:"required,max=100,slug"`
	Content     string `json:"content" validate:"required"`
	Keywords    string `json:"keywords" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=155"`
	Spot        string `json:"spot" validate:"required,max=75"`
	Cover       string `json:"cover" validate:"required,url"`
	Published   bool   `json:"published" validate:"required"`
	CategoryId  int    `json:"categoryId" validate:"required"`
}

type RequestBlogPartial struct {
	Title       string `json:"title" validate:"max=100"`
	Content     string `json:"content"`
	Keywords    string `json:"keywords" validate:"max=255"`
	Description string `json:"description" validate:"max=155"`
	Spot        string `json:"spot" validate:"max=75"`
	Cover       string `json:"cover" validate:"url"`
	Published   bool   `json:"published"`
	CategoryId  int    `json:"categoryId"`
}

type ResponseBlogCard struct {
	Slug         string    `json:"slug"`
	Title        string    `json:"title"`
	Spot         string    `json:"spot"`
	Cover        string    `json:"cover"`
	ReadTime     int       `json:"readTime"`
	ReadCount    int       `json:"readCount"`
	PublishedAt  time.Time `json:"publishedAt"`
	CategoryName string    `json:"categoryName"`
	CategorySlug string    `json:"categorySlug"`
}

type ResponseBlog struct {
	Slug         string    `json:"slug"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Keywords     string    `json:"keywords"`
	Description  string    `json:"description"`
	Spot         string    `json:"spot"`
	Cover        string    `json:"cover"`
	PublishedAt  time.Time `json:"publishedAt"`
	CategoryName string    `json:"categoryName"`
	CategorySlug string    `json:"categorySlug"`
}

func (blog *RequestBlog) CalculateReadTime() int {
	return len(strings.Split(blog.Content, " ")) / wordsPerMinute
}

func (blog *RequestBlog) CalculatePublishedAt() time.Time {
	if blog.Published {
		return time.Now()
	}
	return time.Time{}
}
