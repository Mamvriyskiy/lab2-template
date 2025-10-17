package models

import (
	"github.com/google/uuid"
	"time"
)

type BookReservationResponse struct {
	ReservationUid uuid.UUID         `json:"reservation_uid"`
	Status         ReservationStatus `json:"status"`
	StartDate      time.Time         `json:"start_date"`
	TillDate       time.Time         `json:"till_date"`
	Book           BookInfo          `json:"book"`
	Lib            LibraryResponse   `json:"lib"`
}

type ReservationStatus string

const (
	Rented   ReservationStatus = "RENTED"
	Returned ReservationStatus = "RETURNED"
	Expired  ReservationStatus = "EXPIRED"
)
