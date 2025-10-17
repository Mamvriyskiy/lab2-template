package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reservation/internal"
	"reservation/models"
	"reservation/proto/reservation"
)

type GRPCHandler struct {
	usecase internal.ResUsecase
}

func NewGRPCHandler(usecase internal.ResUsecase) *GRPCHandler {
	return &GRPCHandler{usecase: usecase}
}

func (h *GRPCHandler) FetchReservations(ctx context.Context,
	req *reservation.ReservationFetchRequest) (*reservation.ReservationFetchResponse, error) {
	name := req.Name
	reservations, status := h.usecase.GetReservationsInfo(name)
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", status))
	}
	response := &reservation.ReservationFetchResponse{}
	responseItem := []*reservation.ReservationFetchResponseItem{}

	for _, res := range reservations {
		item := reservation.ReservationFetchResponseItem{}
		item.Status = reservation.ReservationFetchResponseItem_Status(reservation.ReservationFetchResponseItem_Status_value[string(res.Status)])
		item.ReservationUid = res.ReservationUid.String()
		item.StartDate = timestamppb.New(res.StartDate)
		item.TillDate = timestamppb.New(res.TillDate)
		item.LibraryUid = res.Lib.LibraryUid.String()
		item.BookUid = res.Book.BookUid.String()
		responseItem = append(responseItem, &item)
	}
	response.Items = responseItem
	return response, nil
}

func (h *GRPCHandler) TakeBook(ctx context.Context, req *reservation.TakeBookRequest) (*reservation.TakeBookResponse, error) {
	name := req.Name
	bookUidStr := req.BookUid
	libUidStr := req.LibraryUid
	tillDate := req.TillDate.AsTime()

	libUid, err := uuid.Parse(libUidStr)
	if err != nil {
		return nil, err
	}
	bookUid, err := uuid.Parse(bookUidStr)
	if err != nil {
		return nil, err
	}
	usecaseResult, status := h.usecase.TakeBook(name, models.TakeBookRequest{
		BookUid:    bookUid,
		LibraryUid: libUid,
		TillDate:   tillDate})
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", status))
	}
	response := &reservation.TakeBookResponse{}
	response.ReservationUid = usecaseResult.ReservationUid.String()
	response.StartDate = timestamppb.New(usecaseResult.StartDate)
	response.TillDate = timestamppb.New(usecaseResult.TillDate)
	response.Status = reservation.TakeBookResponse_Status(reservation.TakeBookResponse_Status_value[string(usecaseResult.Status)])
	response.LibraryUid = usecaseResult.Library.LibraryUid.String()
	response.BookUid = usecaseResult.Book.BookUid.String()
	return response, nil
}

func (h *GRPCHandler) ReturnBook(ctx context.Context, req *reservation.ReturnBookRequest) (*reservation.ReturnBookResponse, error) {
	reservationUid := req.ReservationUid
	name := req.Name
	condition := req.Condition.String()
	date := req.Date.AsTime()
	uid, err := uuid.Parse(reservationUid)
	if err != nil {
		return nil, err
	}
	retReq := models.ReturnBookRequest{Date: date, Condition: models.BookCondition(condition)}
	status := h.usecase.ReturnBook(uid, name, retReq)
	if status != models.Deleted {
		return &reservation.ReturnBookResponse{Ok: false}, errors.New(fmt.Sprintf("%d", status))
	}
	return &reservation.ReturnBookResponse{Ok: true}, nil
}

func (h *GRPCHandler) GetReservation(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.ReservationFetchResponseItem, error) {
	resUidStr := req.ResUid
	resUid, err := uuid.Parse(resUidStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	res, st := h.usecase.GetReservation(resUid)
	if st != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", models.InternalError))
	}
	item := &reservation.ReservationFetchResponseItem{
		ReservationUid: res.ReservationUid.String(),
		LibraryUid:     res.Lib.LibraryUid.String(),
		BookUid:        res.Book.BookUid.String(),
		Status:         reservation.ReservationFetchResponseItem_Status(reservation.ReservationFetchResponseItem_Status_value[string(res.Status)]),
		TillDate:       timestamppb.New(res.TillDate),
		StartDate:      timestamppb.New(res.StartDate),
	}
	return item, err
}
