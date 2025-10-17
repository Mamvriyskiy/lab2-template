package models

import (
	"github.com/google/uuid"
	"time"
)

type BookReservationResponse struct {
	ReservationUid uuid.UUID         `json:"reservationUid"`
	Status         ReservationStatus `json:"status"`
	StartDate      time.Time         `json:"startDate"`
	TillDate       time.Time         `json:"tillDate"`
	Book           BookInfo          `json:"book"`
	Lib            LibraryResponse   `json:"lib"`
}

type BookReservationResponseParsed struct {
	ReservationUid uuid.UUID         `json:"reservationUid"`
	Status         ReservationStatus `json:"status"`
	StartDate      string            `json:"startDate"`
	TillDate       string            `json:"tillDate"`
	Book           BookInfo          `json:"book"`
	Lib            LibraryResponse   `json:"library"`
}

type ReservationStatus string

const (
	Rented   ReservationStatus = "RENTED"
	Returned ReservationStatus = "RETURNED"
	Expired  ReservationStatus = "EXPIRED"
)
