package usecase

import (
	"gateway/internal"
	models2 "gateway/models"
	"github.com/google/uuid"
	"time"
)

type GatewayUsecase struct {
	client internal.ApiClient
}

func NewGatewayUsecase(client internal.ApiClient) *GatewayUsecase {
	return &GatewayUsecase{
		client: client,
	}
}

func (u *GatewayUsecase) GetLibList(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode) {
	return u.client.GetLibraries(page, size, city)
}
func (u *GatewayUsecase) GetBookList(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode) {
	return u.client.GetBooks(page, size, showAll, LibUid)
}
func (u *GatewayUsecase) GetReservationInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode) {
	reservations, status := u.client.GetReservations(name)
	if status != models2.OK {
		return nil, status
	}

	for i := range reservations {
		bookPag, st := u.client.GetBook(reservations[i].Book.BookUid)
		if st != models2.OK {
			return nil, models2.InternalError
		}
		reservations[i].Book.Name = bookPag.Name
		reservations[i].Book.Author = bookPag.Author
		reservations[i].Book.Genre = bookPag.Genre

		libPag, st2 := u.client.GetLibrary(reservations[i].Lib.LibraryUid)
		if st2 != models2.OK {
			return nil, models2.InternalError
		}
		reservations[i].Lib.City = libPag.City
		reservations[i].Lib.Address = libPag.Address
		reservations[i].Lib.Name = libPag.Name
	}

	return reservations, models2.OK
}
func (u *GatewayUsecase) TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) {
	reservations, st := u.client.GetReservations(name)
	count := 0
	for _, el := range reservations {
		if el.Status == models2.Rented {
			count++
		}
	}
	rating, st := u.client.GetRating(name)
	if count*10 > int(rating.Stars) {
		return models2.TakeBookResponse{}, models2.Forbidden
	}
	resp, st := u.client.TakeBook(name, req)
	if st != models2.OK {
		return models2.TakeBookResponse{}, st
	}
	bookInfo, st := u.client.GetBook(resp.Book.BookUid)
	if st != models2.OK {
		return models2.TakeBookResponse{}, st
	}
	libInfo, st := u.client.GetLibrary(resp.Library.LibraryUid)
	if st != models2.OK {
		return models2.TakeBookResponse{}, st
	}
	resp.Book.Name = bookInfo.Name
	resp.Book.Author = bookInfo.Author
	resp.Book.Genre = bookInfo.Genre
	resp.Library.Name = libInfo.Name
	resp.Library.City = libInfo.City
	resp.Library.Address = libInfo.Address

	ratingInfo, st := u.client.GetRating(name)
	if st != models2.OK {
		return models2.TakeBookResponse{}, st
	}
	st = u.client.UpdateBooksCount(resp.Book.BookUid, -1)
	if st != models2.OK {
		return models2.TakeBookResponse{}, st
	}
	resp.Rating.Stars = ratingInfo.Stars
	return resp, models2.OK
}
func (u *GatewayUsecase) ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode {
	reservation, st := u.client.GetReservation(resUid)
	if st != models2.OK {
		return models2.NotFound
	}
	book, st := u.client.GetBook(reservation.Book.BookUid)
	if st != models2.OK {
		return models2.BadRequest
	}
	if req.Date.Unix() > time.Now().Unix() || book.Condition != req.Condition {
		u.client.UpdateRating(name, -10)
	} else {
		u.client.UpdateRating(name, 1)
	}
	st = u.client.UpdateBooksCount(book.BookUid, 1)
	if st != models2.OK {
		return st
	}
	return u.client.ReturnBook(resUid, name, req)
}
func (u *GatewayUsecase) GetRating(name string) (models2.UserRatingResponse, models2.StatusCode) {
	rating, st := u.client.GetRating(name)
	return rating, st
}
