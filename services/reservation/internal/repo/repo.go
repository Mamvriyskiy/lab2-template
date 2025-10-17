package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	models2 "reservation/models"
	"time"
)

const (
	SelectReservation = "SELECT r.reservation_uid, r.status, r.start_date, r.till_date, " +
		"r.book_uid, " +
		"r.library_uid " +
		"FROM reservations.reservation r " +
		"WHERE r.username=$1"

	InsertReservation = "INSERT INTO reservations.reservation " +
		"(reservation_uid, username, book_uid, library_uid, status, start_date, till_date )" +
		"VALUES($1,$2,$3,$4,$5,$6,$7)"

	UpdateReservation = "UPDATE reservations.reservation SET status=$1 WHERE reservation_uid=$2"

	GetReservation = "SELECT r.reservation_uid, r.status, r.start_date, r.till_date, " +
		"r.book_uid, " +
		"r.library_uid " +
		"FROM reservations.reservation r " +
		"WHERE r.reservation_uid=$1"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) ReservationsInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode) {
	var fullInfo []models2.BookReservationResponse
	res, err := r.conn.Query(context.Background(), SelectReservation, name)
	if err != nil {
		return []models2.BookReservationResponse{}, models2.InternalError
	}
	for res.Next() {
		reservation := models2.BookReservationResponse{}
		res.Scan(&reservation.ReservationUid, &reservation.Status, &reservation.StartDate, &reservation.TillDate,
			&reservation.Book.BookUid, &reservation.Lib.LibraryUid)
		fullInfo = append(fullInfo, reservation)
	}
	return fullInfo, models2.OK
}
func (r *Repo) ReserveBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) {
	resUid := uuid.New()
	start_date := time.Now()
	_, err := r.conn.Exec(context.Background(), InsertReservation, resUid, name,
		req.BookUid, req.LibraryUid, models2.Rented, start_date, req.TillDate)
	if err != nil {
		return models2.TakeBookResponse{}, models2.BadRequest
	}

	reservation := models2.TakeBookResponse{
		ReservationUid: resUid,
		TillDate:       req.TillDate,
		StartDate:      start_date,
	}
	reservation.Book.BookUid = req.BookUid
	reservation.Library.LibraryUid = req.LibraryUid
	reservation.Status = models2.Rented

	return reservation, models2.OK
}
func (r *Repo) ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode {
	var status models2.ReservationStatus
	if req.Date.Unix() > time.Now().Unix() {
		status = models2.Expired
	} else {
		status = models2.Returned
	}

	_, err := r.conn.Exec(context.Background(), UpdateReservation, status, resUid.String())
	if err != nil {
		return models2.NotFound
	}
	return models2.Deleted
}

func (r *Repo) GetReservation(resUid uuid.UUID) (models2.BookReservationResponse, models2.StatusCode) {
	res := r.conn.QueryRow(context.Background(), GetReservation, resUid.String())
	reservation := models2.BookReservationResponse{}
	reservation.ReservationUid = resUid
	err := res.Scan(&reservation.ReservationUid, &reservation.Status, &reservation.StartDate, &reservation.TillDate,
		&reservation.Book.BookUid, &reservation.Lib.LibraryUid)
	if err != nil {
		return models2.BookReservationResponse{}, models2.InternalError
	}
	return reservation, models2.OK
}
