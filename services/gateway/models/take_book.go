package models

import (
	"github.com/google/uuid"
	"time"
)

type TakeBookRequest struct {
	BookUid    uuid.UUID `json:"bookUid"`
	LibraryUid uuid.UUID `json:"libraryUid"`
	TillDate   time.Time `json:"tillDate"`
}

type TakeBookRequestPreParsed struct {
	BookUid    uuid.UUID `json:"bookUid"`
	LibraryUid uuid.UUID `json:"libraryUid"`
	TillDate   string    `json:"tillDate"`
}

type TakeBookResponse struct {
	ReservationUid uuid.UUID          `json:"reservationUid"`
	Status         ReservationStatus  `json:"status"`
	StartDate      time.Time          `json:"startDate"`
	TillDate       time.Time          `json:"tillDate"`
	Book           BookInfo           `json:"book"`
	Library        LibraryResponse    `json:"library"`
	Rating         UserRatingResponse `json:"rating"`
}

type TakeBookResponseParsedTime struct {
	ReservationUid uuid.UUID          `json:"reservationUid"`
	Status         ReservationStatus  `json:"status"`
	StartDate      string             `json:"startDate"`
	TillDate       string             `json:"tillDate"`
	Book           BookInfo           `json:"book"`
	Library        LibraryResponse    `json:"library"`
	Rating         UserRatingResponse `json:"rating"`
}
