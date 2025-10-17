package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"library/models"
)

const (
	// TODO count
	GetBooksQuery = "SELECT b.id, b.book_uid, b.name, b.author, b.genre, b.condition, lb.available_count, COUNT(*) FROM libraries.books b " +
		"INNER JOIN libraries.library_books lb ON lb.book_id=b.id " + "INNER JOIN libraries.library l ON l.id=lb.library_id " +
		"WHERE l.library_uid=$1 AND lb.available_count>$2 " +
		"GROUP BY b.id, lb.available_count " +
		"LIMIT $3 OFFSET $4"
	GetLibsQuery    = "SELECT id, library_uid, name, city, address, COUNT(*) FROM libraries.library WHERE city=$1 GROUP BY(library.id) LIMIT $2 OFFSET $3"
	GetOneBookQuery = "SELECT name, author, genre, condition FROM libraries.books WHERE book_uid=$1 LIMIT 1"
	GetOneLibQuery  = "SELECT name, city, address FROM libraries.library WHERE library_uid=$1 LIMIT 1"
	UpdateBookQuery = "WITH b AS (SELECT id FROM libraries.books WHERE book_uid=$2)" +
		"UPDATE libraries.library_books SET available_count=available_count+$1 FROM b WHERE book_id=b.id"
)

type LibRepo struct {
	conn *pgxpool.Pool
}

func NewLibRepo(conn *pgxpool.Pool) *LibRepo {
	return &LibRepo{conn: conn}
}

func (lr *LibRepo) GetLibraries(page, size int64, city string) ([]models.LibraryResponse, int64, models.StatusCode) {
	rows, err := lr.conn.Query(context.Background(), GetLibsQuery, city, size, page-1)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		libs  []models.LibraryResponse
		count int64
	)
	for rows.Next() {
		var lib models.LibraryResponse
		err := rows.Scan(&lib.Id, &lib.LibraryUid, &lib.Name, &lib.City, &lib.Address, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		libs = append(libs, lib)
	}
	return libs, count, models.OK
}

func (lr *LibRepo) GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models.LibraryBookResponse, int64, models.StatusCode) {
	includeZeroCountOfBooks := 0
	if showAll {
		includeZeroCountOfBooks = -1
	}

	rows, err := lr.conn.Query(context.Background(), GetBooksQuery, LibUid, includeZeroCountOfBooks, size, page-1)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		books []models.LibraryBookResponse
		count int64
	)
	for rows.Next() {
		var book models.LibraryBookResponse
		err := rows.Scan(&book.Id, &book.BookId, &book.Name, &book.Author,
			&book.Genre, &book.Condition, &book.AvailableCount, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		books = append(books, book)
	}
	return books, count, models.OK
}

func (lr *LibRepo) GetBook(bookUid uuid.UUID) (models.BookInfo, models.StatusCode) {
	row := lr.conn.QueryRow(context.Background(), GetOneBookQuery, bookUid.String())
	book := models.BookInfo{BookUid: bookUid}
	err := row.Scan(&book.Name, &book.Author, &book.Genre, &book.Condition)
	if err != nil {
		return models.BookInfo{}, models.InternalError
	}
	return book, models.OK
}

func (lr *LibRepo) GetLib(libUid uuid.UUID) (models.LibraryResponse, models.StatusCode) {
	row := lr.conn.QueryRow(context.Background(), GetOneLibQuery, libUid.String())
	lib := models.LibraryResponse{LibraryUid: libUid}
	err := row.Scan(&lib.Name, &lib.City, &lib.Address)
	if err != nil {
		return models.LibraryResponse{}, models.InternalError
	}
	return lib, models.OK
}

func (lr *LibRepo) UpdateBookCount(bookUid uuid.UUID, num int) models.StatusCode {
	_, err := lr.conn.Exec(context.Background(), UpdateBookQuery, num, bookUid.String())
	if err != nil {
		return models.InternalError
	}
	return models.OK
}
