package models

import (
	"github.com/google/uuid"
	"time"
)

type TakeBookRequest struct {
	BookUid    uuid.UUID `json:"book_uid"`
	LibraryUid uuid.UUID `json:"library_uid"`
	TillDate   time.Time `json:"till_date"`
}

type TakeBookResponse struct {
	ReservationUid uuid.UUID          `json:"reservation_uid"`
	Status         ReservationStatus  `json:"status"`
	StartDate      time.Time          `json:"start_date"`
	TillDate       time.Time          `json:"till_date"`
	Book           BookInfo           `json:"book"`
	Library        LibraryResponse    `json:"library"`
	Rating         UserRatingResponse `json:"rating"`
}
