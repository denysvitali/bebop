package store

import (
	"time"
	"unicode/utf8"
)

// Category is a general theme of topics.
type Category struct {
	ID          int64     `json:"id"`
	ParentID    int64     `json:"parentId"`
	AuthorID    int64     `json:"authorId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	LastTopicAt time.Time `json:"lastTopicAt"`
	TopicCount  int       `json:"topicCount"`
}

const (
	categoryTitleMinLen       = 1
	categoryTitleMaxLen       = 100
	categoryDescriptionMinLen = 0
	categoryDescriptionMaxLen = 100
)

// ValidCategoryTitle checks if a category title is valid.
func ValidCategoryTitle(title string) bool {
	if !utf8.ValidString(title) {
		return false
	}

	length := utf8.RuneCountInString(title)
	if !(categoryTitleMinLen <= length && length <= categoryTitleMaxLen) {
		return false
	}

	return true
}

// ValidCategoryDescription checks if a category description is valid.
func ValidCategoryDescription(description string) bool {
	if !utf8.ValidString(description) {
		return false
	}

	length := utf8.RuneCountInString(description)
	if !(categoryDescriptionMinLen <= length && length <= categoryDescriptionMaxLen) {
		return false
	}

	return true
}
