package models

import "github.com/google/uuid"

type LibraryBookResponse struct {
	Id             uint          `json:"id,omitempty"`
	BookId         uuid.UUID     `json:"book_uid"`
	Name           string        `json:"name"`
	Author         string        `json:"author"`
	Genre          string        `json:"genre"`
	Condition      BookCondition `json:"condition"`
	AvailableCount int32         `json:"available_count"`
}

type BookCondition string

const (
	Excellent BookCondition = "EXCELLENT"
	Good      BookCondition = "GOOD"
	Bad       BookCondition = "BAD"
)
