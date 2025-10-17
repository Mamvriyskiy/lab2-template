package usecase

import (
	"github.com/google/uuid"
	"reservation/internal"
	models2 "reservation/models"
)

type Usecase struct {
	repo internal.ResRepo
}

func NewUsecase(repo internal.ResRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetReservationsInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode) {
	return u.repo.ReservationsInfo(name)
}

func (u *Usecase) TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) {
	return u.repo.ReserveBook(name, req)
}

func (u *Usecase) ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode {
	return u.repo.ReturnBook(resUid, name, req)
}

func (u *Usecase) GetReservation(resUid uuid.UUID) (models2.BookReservationResponse, models2.StatusCode) {
	return u.repo.GetReservation(resUid)
}
