package models

import "github.com/google/uuid"

type BookInfo struct {
	BookUid uuid.UUID `json:"book_uid"`
	Name    string    `json:"name"`
	Author  string    `json:"author"`
	Genre   string    `json:"genre"`
}
