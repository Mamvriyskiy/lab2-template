package delivery

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"library/internal"
	models2 "library/models"
	"library/proto/library"
	"testing"
)

func TestGRPCHandler_FetchBooks(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockLibraryUsecase(ctl)
	page, size := int64(1), int64(1)
	showAll := true
	libUid := uuid.New()
	bookUid := uuid.New()
	req := &library.BookRequest{
		Page:       page,
		Size:       size,
		ShowAll:    showAll,
		LibraryUid: libUid.String(),
	}
	bookItem := &library.ItemBook{
		BookUid:        bookUid.String(),
		Name:           "test",
		Condition:      1,
		Author:         "test",
		Genre:          "test",
		AvailableCount: 1,
	}
	usecaseRes := []models2.LibraryBookPaginationResponse{
		models2.LibraryBookPaginationResponse{
			Page:          page,
			PageSize:      size,
			TotalElements: 1,
			Items: []models2.LibraryBookResponse{
				models2.LibraryBookResponse{
					BookId:         bookUid,
					Name:           bookItem.Name,
					Author:         bookItem.Author,
					Genre:          bookItem.Genre,
					Condition:      models2.BookCondition(bookItem.Condition),
					AvailableCount: bookItem.AvailableCount,
				},
			},
		},
	}
	mockUsecase.EXPECT().GetBooksList(page, size, showAll, libUid).Return(usecaseRes, models2.OK)
	handler := NewGRPCHandler(mockUsecase)

	testResult, err := handler.FetchBooks(context.Background(), req)
	if err != nil || testResult.Items[0].Item[0].BookUid != bookUid.String() {
		t.Error("Incorrect test result")
	}
}

func TestGRPCHandler_FetchLibs(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUsecase := internal.NewMockLibraryUsecase(ctl)
	page, size := int64(1), int64(1)
	city := "test"
	libUid := uuid.New()
	req := &library.LibraryRequest{
		Page: page,
		Size: size,
		City: city,
	}
	libItem := &library.ItemLibrary{
		LibraryUid: libUid.String(),
		Name:       "test",
		Address:    city,
		City:       city,
	}
	usecaseRes := []models2.LibraryPaginationResponse{
		models2.LibraryPaginationResponse{
			Page:          page,
			PageSize:      size,
			TotalElements: 1,
			Items: []models2.LibraryResponse{
				models2.LibraryResponse{
					LibraryUid: libUid,
					Name:       libItem.Name,
					City:       city,
					Address:    city,
				},
			},
		},
	}
	mockUsecase.EXPECT().GetLibrariesList(page, size, city).Return(usecaseRes, models2.OK)
	handler := NewGRPCHandler(mockUsecase)

	testResult, err := handler.FetchLibs(context.Background(), req)
	if err != nil || testResult.Items[0].Item[0].LibraryUid != libUid.String() {
		t.Error("Incorrect test result")
	}
}
