package models

import "github.com/google/uuid"

type LibraryResponse struct {
	Id         uint      `json:"id,omitempty"`
	LibraryUid uuid.UUID `json:"library_uid"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	Address    string    `json:"address"`
}
