package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"library/internal"
	"library/models"
	"library/proto/library"
	"log"
)

type GRPCHandler struct {
	usecase internal.LibraryUsecase
}

func NewGRPCHandler(usecase internal.LibraryUsecase) *GRPCHandler {
	return &GRPCHandler{usecase: usecase}
}

func (h *GRPCHandler) FetchLibs(ctx context.Context, req *library.LibraryRequest) (*library.LibraryResponse, error) {
	page := req.Page
	size := req.Size
	city := req.City
	libs, status := h.usecase.GetLibrariesList(page, size, city)
	log.Println("HERE")
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", status))
	}
	response := &library.LibraryResponse{}
	responseItem := library.LibraryResponseItem{Page: page, Size: size, TotalElements: libs[0].TotalElements}

	for i := range libs {
		for _, el := range libs[i].Items {
			itemLib := library.ItemLibrary{LibraryUid: el.LibraryUid.String(), Name: el.Name,
				Address: el.Address, City: el.City}
			responseItem.Item = append(responseItem.Item, &itemLib)
		}
	}
	response.Items = append(response.Items, &responseItem)
	return response, nil
}

func (h *GRPCHandler) FetchBooks(ctx context.Context, req *library.BookRequest) (*library.BookResponse, error) {
	page := req.Page
	size := req.Size
	libUid := req.LibraryUid
	showAll := req.ShowAll

	uid, err := uuid.Parse(libUid)
	if err != nil {
		return nil, err
	}
	libs, status := h.usecase.GetBooksList(page, size, showAll, uid)
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", models.OK))
	}

	response := &library.BookResponse{}
	responseItem := library.BookResponseItem{Page: page, Size: size, TotalElements: libs[0].TotalElements}
	for i := range libs {
		for _, el := range libs[i].Items {
			itemBook := library.ItemBook{Name: el.Name, BookUid: el.BookId.String(), Author: el.Author,
				Genre: el.Genre, AvailableCount: el.AvailableCount, Condition: library.ItemBook_Condition(library.ItemBook_Condition_value[string(el.Condition)])}
			responseItem.Item = append(responseItem.Item, &itemBook)
		}
	}
	response.Items = append(response.Items, &responseItem)
	return response, nil
}

func (h *GRPCHandler) GetBook(ctx context.Context, req *library.GetOneBookRequest) (*library.ItemBook, error) {
	response := &library.ItemBook{}
	uid, err := uuid.Parse(req.BookUid)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	book, st := h.usecase.GetBook(uid)
	if st != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	response.BookUid = book.BookUid.String()
	response.Name = book.Name
	response.Genre = book.Genre
	response.Author = book.Author
	response.Condition = library.ItemBook_Condition(library.ItemBook_Condition_value[string(book.Condition)])
	return response, nil
}

func (h *GRPCHandler) GetLibrary(ctx context.Context, req *library.GetOneLibRequest) (*library.ItemLibrary, error) {
	response := &library.ItemLibrary{}
	uid, err := uuid.Parse(req.LibUid)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	book, st := h.usecase.GetLib(uid)
	if st != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	response.LibraryUid = book.LibraryUid.String()
	response.Name = book.Name
	response.City = book.City
	response.Address = book.Address
	return response, nil
}

func (h *GRPCHandler) UpdateBookCount(ctx context.Context, req *library.UpdateBookCountRequest) (*library.UpdateBookCountResponse, error) {
	uid, err := uuid.Parse(req.BookUid)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d", models.BadRequest))
	}
	st := h.usecase.UpdateBookCount(uid, int(req.Num))
	if st != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", models.InternalError))
	}
	return &library.UpdateBookCountResponse{Ok: true}, nil
}
