package delivery

import (
	"context"
	models2 "gateway/models"
	"gateway/proto/library"
	"gateway/proto/rating"
	"gateway/proto/reservation"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	RatingServiceClient      rating.RatingServiceClient
	ReservationServiceClient reservation.ReservationServiceClient
	LibraryServiceClient     library.LibraryServiceClient
}

func NewGRPCClient(Rating, Reservation, Library grpc.ClientConnInterface) *Client {
	return &Client{
		RatingServiceClient:      rating.NewRatingServiceClient(Rating),
		ReservationServiceClient: reservation.NewReservationServiceClient(Reservation),
		LibraryServiceClient:     library.NewLibraryServiceClient(Library),
	}
}

func (cl *Client) GetBook(bookId uuid.UUID) (models2.BookInfo, models2.StatusCode) {
	request := &library.GetOneBookRequest{BookUid: bookId.String()}
	response, err := cl.LibraryServiceClient.GetBook(context.Background(), request)
	if err != nil {
		return models2.BookInfo{}, models2.InternalError
	}
	condition := library.ItemBook_Condition_name[int32(response.Condition)]
	book := models2.BookInfo{BookUid: bookId, Name: response.Name, Author: response.Author, Genre: response.Genre, Condition: models2.BookCondition(condition)}
	return book, models2.OK
}

func (cl *Client) GetLibrary(libId uuid.UUID) (models2.LibraryResponse, models2.StatusCode) {
	request := &library.GetOneLibRequest{LibUid: libId.String()}
	response, err := cl.LibraryServiceClient.GetLibrary(context.Background(), request)
	if err != nil {
		return models2.LibraryResponse{}, models2.InternalError
	}
	lib := models2.LibraryResponse{LibraryUid: libId, Name: response.Name, City: response.City, Address: response.Address}
	return lib, models2.OK
}

func (cl *Client) GetLibraries(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode) {
	request := &library.LibraryRequest{Page: page, Size: size, City: city}
	response, err := cl.LibraryServiceClient.FetchLibs(context.Background(), request)
	if err != nil {
		return nil, models2.InternalError
	}
	libs := []models2.LibraryPaginationResponse{}
	for _, el := range response.Items {
		libEl := models2.LibraryPaginationResponse{}
		libEl.TotalElements = el.TotalElements
		libEl.Page = el.Page
		libEl.PageSize = el.Size
		for _, item := range el.Item {
			uid, err := uuid.Parse(item.LibraryUid)
			if err != nil {
				return nil, models2.BadRequest
			}
			libEl.Items = append(libEl.Items, models2.LibraryResponse{
				LibraryUid: uid,
				Name:       item.Name,
				Address:    item.Address,
				City:       item.City,
			})
			libs = append(libs, libEl)
		}
	}
	return libs, models2.OK
}
func (cl *Client) GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode) {
	request := &library.BookRequest{Page: page, Size: size, ShowAll: showAll, LibraryUid: LibUid.String()}
	response, err := cl.LibraryServiceClient.FetchBooks(context.Background(), request)
	if err != nil {
		return nil, models2.InternalError
	}
	libs := []models2.LibraryBookPaginationResponse{}
	for _, el := range response.Items {
		libEl := models2.LibraryBookPaginationResponse{}
		libEl.TotalElements = el.TotalElements
		libEl.Page = el.Page
		libEl.PageSize = el.Size
		for _, item := range el.Item {
			uid, err := uuid.Parse(item.BookUid)
			if err != nil {
				return nil, models2.BadRequest
			}
			libEl.Items = append(libEl.Items, models2.LibraryBookResponse{
				BookId:         uid,
				Name:           item.Name,
				Author:         item.Author,
				Genre:          item.Genre,
				Condition:      models2.BookCondition(item.Condition.String()),
				AvailableCount: item.AvailableCount,
			})
			libs = append(libs, libEl)
		}
	}
	return libs, models2.OK
}
func (cl *Client) GetReservations(name string) ([]models2.BookReservationResponse, models2.StatusCode) {
	request := &reservation.ReservationFetchRequest{Name: name}
	response, err := cl.ReservationServiceClient.FetchReservations(context.Background(), request)
	if err != nil {
		return nil, models2.InternalError
	}
	resp := []models2.BookReservationResponse{}
	for _, el := range response.Items {
		uid, err := uuid.Parse(el.ReservationUid)
		if err != nil {
			return nil, models2.BadRequest
		}
		bookUid, err := uuid.Parse(el.BookUid)
		if err != nil {
			return nil, models2.BadRequest
		}
		libUid, err := uuid.Parse(el.LibraryUid)
		if err != nil {
			return nil, models2.BadRequest
		}
		resp = append(resp, models2.BookReservationResponse{
			ReservationUid: uid,
			Status:         models2.ReservationStatus(el.Status.String()),
			StartDate:      el.StartDate.AsTime(),
			TillDate:       el.TillDate.AsTime(),
			Book:           models2.BookInfo{BookUid: bookUid},
			Lib:            models2.LibraryResponse{LibraryUid: libUid},
		})
	}
	return resp, models2.OK
}

func (cl *Client) TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) {
	request := &reservation.TakeBookRequest{
		Name:       name,
		TillDate:   timestamppb.New(req.TillDate),
		BookUid:    req.BookUid.String(),
		LibraryUid: req.LibraryUid.String(),
	}
	response, err := cl.ReservationServiceClient.TakeBook(context.Background(), request)
	if err != nil {
		return models2.TakeBookResponse{}, models2.InternalError
	}
	resUid, err := uuid.Parse(response.ReservationUid)
	if err != nil {
		return models2.TakeBookResponse{}, models2.BadRequest
	}
	bookUid, err := uuid.Parse(response.BookUid)
	if err != nil {
		return models2.TakeBookResponse{}, models2.BadRequest
	}
	libUid, err := uuid.Parse(response.LibraryUid)
	if err != nil {
		return models2.TakeBookResponse{}, models2.BadRequest
	}
	resp := models2.TakeBookResponse{
		ReservationUid: resUid,
		Library:        models2.LibraryResponse{LibraryUid: libUid},
		Book:           models2.BookInfo{BookUid: bookUid},
		Status:         models2.ReservationStatus(response.Status.String()),
		StartDate:      response.StartDate.AsTime(),
		TillDate:       response.TillDate.AsTime(),
	}
	return resp, 0
}
func (cl *Client) ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode {
	request := &reservation.ReturnBookRequest{
		ReservationUid: resUid.String(),
		Name:           name,
		Condition:      reservation.ReturnBookRequest_Condition(reservation.ReturnBookRequest_Condition_value[string(req.Condition)]),
		Date:           timestamppb.New(req.Date),
	}
	response, err := cl.ReservationServiceClient.ReturnBook(context.Background(), request)
	if err != nil {
		return models2.InternalError
	}
	if !response.Ok {
		return models2.NotFound
	}
	return models2.Deleted
}
func (cl *Client) GetRating(name string) (models2.UserRatingResponse, models2.StatusCode) {
	request := &rating.RatingRequest{Name: name}
	response, err := cl.RatingServiceClient.GetRating(context.Background(), request)
	if err != nil {
		return models2.UserRatingResponse{}, models2.BadRequest
	}
	return models2.UserRatingResponse{Stars: uint(response.Stars)}, 0
}

func (cl *Client) UpdateRating(name string, num int32) models2.StatusCode {
	request := &rating.RatingUpdateRequest{Name: name, Add: num}
	response, err := cl.RatingServiceClient.RatingUpdate(context.Background(), request)
	if err != nil || !response.Ok {
		return models2.InternalError
	}
	return models2.OK
}

func (cl *Client) GetReservation(resUid uuid.UUID) (models2.BookReservationResponse, models2.StatusCode) {
	request := &reservation.GetReservationRequest{ResUid: resUid.String()}
	res, err := cl.ReservationServiceClient.GetReservation(context.Background(), request)
	if err != nil {
		return models2.BookReservationResponse{}, models2.InternalError
	}
	uid, err := uuid.Parse(res.ReservationUid)
	if err != nil {
		return models2.BookReservationResponse{}, models2.BadRequest
	}
	bookUid, err := uuid.Parse(res.BookUid)
	if err != nil {
		return models2.BookReservationResponse{}, models2.InternalError
	}
	libUid, err := uuid.Parse(res.LibraryUid)
	if err != nil {
		return models2.BookReservationResponse{}, models2.InternalError
	}

	ret := models2.BookReservationResponse{
		ReservationUid: uid,
		Status:         models2.ReservationStatus(res.Status.String()),
		StartDate:      res.StartDate.AsTime(),
		TillDate:       res.TillDate.AsTime(),
		Book:           models2.BookInfo{BookUid: bookUid},
		Lib:            models2.LibraryResponse{LibraryUid: libUid},
	}
	return ret, models2.OK
}

func (cl *Client) UpdateBooksCount(bookUid uuid.UUID, num int) models2.StatusCode {
	request := &library.UpdateBookCountRequest{BookUid: bookUid.String(), Num: int32(num)}
	resp, err := cl.LibraryServiceClient.UpdateBookCount(context.Background(), request)
	if err != nil || !resp.Ok {
		return models2.InternalError
	}
	return models2.OK
}
