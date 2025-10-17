package usecase

import (
	"github.com/google/uuid"
	"library/internal"
	"library/models"
)

type LibUsecase struct {
	repo internal.LibraryRepo
}

func NewLibUsecase(repo internal.LibraryRepo) *LibUsecase {
	return &LibUsecase{repo: repo}
}

func (lu *LibUsecase) GetLibrariesList(page, size int64, city string) ([]models.LibraryPaginationResponse, models.StatusCode) {
	libs, count, status := lu.repo.GetLibraries(page, size, city)
	if status != models.OK {
		return nil, status
	}
	answer := models.LibraryPaginationResponse{Page: page, PageSize: size, TotalElements: count, Items: libs}
	return []models.LibraryPaginationResponse{answer}, status
}

func (lu *LibUsecase) GetBooksList(page, size int64, showAll bool, LibUid uuid.UUID) ([]models.LibraryBookPaginationResponse, models.StatusCode) {
	books, count, status := lu.repo.GetBooks(page, size, showAll, LibUid)
	if status != models.OK {
		return nil, status
	}
	answer := models.LibraryBookPaginationResponse{Page: page, PageSize: size, TotalElements: count, Items: books}
	return []models.LibraryBookPaginationResponse{answer}, status
}

func (lu *LibUsecase) GetBook(bookUid uuid.UUID) (models.BookInfo, models.StatusCode) {
	return lu.repo.GetBook(bookUid)
}

func (lu *LibUsecase) GetLib(libUid uuid.UUID) (models.LibraryResponse, models.StatusCode) {
	return lu.repo.GetLib(libUid)
}

func (lu *LibUsecase) UpdateBookCount(bookUid uuid.UUID, num int) models.StatusCode {
	return lu.repo.UpdateBookCount(bookUid, num)
}
