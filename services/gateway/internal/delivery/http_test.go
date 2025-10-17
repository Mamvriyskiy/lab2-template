package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal"
	models2 "gateway/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGatewayHandler_GetBook(t *testing.T) {
	date := "2021-10-10"
	req := models2.TakeBookRequest{}
	req.BookUid = uuid.New()
	req.LibraryUid = uuid.New()
	req.TillDate, _ = time.Parse(layout, date)
	reqPreParsed := models2.TakeBookRequestPreParsed{BookUid: req.BookUid, LibraryUid: req.LibraryUid, TillDate: date}
	name := "TestUser"
	response := models2.TakeBookResponse{ReservationUid: uuid.New()}
	body, _ := json.Marshal(reqPreParsed)

	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	mockUsecase.EXPECT().TakeBook(name, req).Return(response, models2.OK)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader(body))
	r.Header.Set("X-User-Name", name)

	handler := NewGatewayHandler(mockUsecase)
	handler.GetBook(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong result by get book method")
	}
}

func TestGatewayHandler_GetBooks(t *testing.T) {
	page, size := int64(1), int64(1)
	showAll := true
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?page=1&size=1&showAll=true", bytes.NewReader(nil))
	r.URL.Query().Set("showAll", "true")
	r.URL.Query().Set("size", fmt.Sprintf("%d", size))
	r.URL.Query().Set("page", fmt.Sprintf("%d", page))

	libUid := uuid.New()
	r = mux.SetURLVars(r, map[string]string{
		"libraryUid": libUid.String(),
	})
	response := []models2.LibraryBookPaginationResponse{
		models2.LibraryBookPaginationResponse{},
	}
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	mockUsecase.EXPECT().GetBookList(page, size, showAll, libUid).Return(response, models2.OK)

	handler := NewGatewayHandler(mockUsecase)
	handler.GetBooks(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong result by get bookList method")
	}
}

func TestGatewayHandler_GetLibraries(t *testing.T) {
	page, size := int64(1), int64(1)
	city := "Moscow"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?page=1&size=1&city=Moscow", bytes.NewReader(nil))

	response := []models2.LibraryPaginationResponse{
		models2.LibraryPaginationResponse{},
	}
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	mockUsecase.EXPECT().GetLibList(page, size, city).Return(response, models2.OK)

	handler := NewGatewayHandler(mockUsecase)
	handler.GetLibraries(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong result by get bookList method")
	}
}

func TestGatewayHandler_GetRating(t *testing.T) {
	name := "TestUser"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader(nil))
	r.Header.Set("X-User-Name", name)
	response := models2.UserRatingResponse{}
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	mockUsecase.EXPECT().GetRating(name).Return(response, models2.OK)

	handler := NewGatewayHandler(mockUsecase)
	handler.GetRating(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong result by get bookList method")
	}
}

func TestGatewayHandler_GetReservations(t *testing.T) {
	name := "TestUser"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader(nil))
	r.Header.Set("X-User-Name", name)
	resUid := uuid.New()
	status := "RENTED"
	date := time.Now()
	libUid := uuid.New()
	bookUid := uuid.New()
	response := []models2.BookReservationResponse{
		models2.BookReservationResponse{
			ReservationUid: resUid,
			Status:         models2.ReservationStatus(status),
			StartDate:      time.Now(),
			TillDate:       date,
			Book:           models2.BookInfo{BookUid: bookUid},
			Lib:            models2.LibraryResponse{LibraryUid: libUid},
		},
	}
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	mockUsecase.EXPECT().GetReservationInfo(name).Return(response, models2.OK)

	handler := NewGatewayHandler(mockUsecase)
	handler.GetReservations(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong result by get bookList method")
	}
}

func TestGatewayHandler_ReturnBook(t *testing.T) {
	name := "TestUser"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader(nil))
	r.Header.Set("X-User-Name", name)
	//resUid := uuid.New()
	//date, _ := time.Parse(layout, time.Now().String())
	//req := models2.ReturnBookRequest{
	//	Condition: models2.Excellent,
	//	Date: date,
	//}
	//r = mux.SetURLVars(r, map[string]string{
	//	"reservationUid": resUid.String(),
	//})
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockUsecase(ctl)
	//mockUsecase.EXPECT().ReturnBook(resUid, name, req).Return(models2.OK)

	handler := NewGatewayHandler(mockUsecase)
	handler.ReturnBook(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Error("Wrong result by get bookList method")
	}
}
