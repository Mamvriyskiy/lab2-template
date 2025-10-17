package delivery

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reservation/internal"
	"reservation/models"
	"reservation/proto/reservation"
	"testing"
	"time"
)

func TestGRPCHandler_FetchReservations(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockResUsecase(ctl)
	name := "test"
	resUid := uuid.New()
	req := &reservation.ReservationFetchRequest{
		Name: name,
	}
	resList := []models.BookReservationResponse{
		models.BookReservationResponse{
			ReservationUid: resUid,
		},
	}
	mockUsecase.EXPECT().GetReservationsInfo(name).Return(resList, models.OK)
	handler := NewGRPCHandler(mockUsecase)
	result, err := handler.FetchReservations(context.Background(), req)
	if err != nil || result.Items[0].ReservationUid != resUid.String() {
		t.Error("incorrect result")
	}
}

func TestGRPCHandler_ReturnBook(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockResUsecase(ctl)
	name := "test"
	resUid := uuid.New()
	date := time.Now()
	req := &reservation.ReturnBookRequest{
		ReservationUid: resUid.String(),
		Name:           name,
		Date:           timestamppb.New(date),
		Condition:      reservation.ReturnBookRequest_Condition(reservation.ReturnBookRequest_Condition_value["BAD"]),
	}

	reqRet := models.ReturnBookRequest{
		Date:      timestamppb.New(date).AsTime(),
		Condition: models.BookCondition("BAD"),
	}

	mockUsecase.EXPECT().ReturnBook(resUid, name, reqRet).Return(models.Deleted)
	handler := NewGRPCHandler(mockUsecase)
	result, err := handler.ReturnBook(context.Background(), req)
	if err != nil || !result.Ok {
		t.Error("incorrect result")
	}
}

func TestGRPCHandler_TakeBook(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockResUsecase(ctl)
	name := "test"
	resUid := uuid.New()
	libUid := uuid.New()
	bookUid := uuid.New()
	tillDate := time.Now()
	req := &reservation.TakeBookRequest{
		BookUid:    bookUid.String(),
		Name:       name,
		LibraryUid: libUid.String(),
		TillDate:   timestamppb.New(tillDate),
	}
	uscReq := models.TakeBookRequest{
		BookUid:    bookUid,
		LibraryUid: libUid,
		TillDate:   req.TillDate.AsTime(),
	}
	resp := models.TakeBookResponse{
		ReservationUid: resUid,
		Status:         models.Rented,
		StartDate:      tillDate,
		TillDate:       tillDate,
		Book:           models.BookInfo{BookUid: bookUid},
		Library:        models.LibraryResponse{LibraryUid: libUid},
		Rating:         models.UserRatingResponse{Stars: 75},
	}

	mockUsecase.EXPECT().TakeBook(name, uscReq).Return(resp, models.OK)
	handler := NewGRPCHandler(mockUsecase)
	result, err := handler.TakeBook(context.Background(), req)
	if err != nil || result.BookUid != bookUid.String() {
		t.Error("incorrect result")
	}
}
