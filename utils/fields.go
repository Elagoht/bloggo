package utils

import (
	"strings"
	"time"

	"github.com/gosimple/slug"
)

const (
	wordsPerMinute = 180
)

func GetReadTime(content string) int {
	return len(strings.Split(content, " ")) / wordsPerMinute
}

func CalculatePublishedAt(published bool) time.Time {
	if published {
		return time.Now()
	}
	return time.Time{}
}

func GenerateSlug(title string) string {
	return slug.Make(title)
}
