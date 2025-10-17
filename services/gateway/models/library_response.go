package models

import "github.com/google/uuid"

type LibraryResponse struct {
	Id         uint      `json:"id,omitempty"`
	LibraryUid uuid.UUID `json:"libraryUid"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
}
