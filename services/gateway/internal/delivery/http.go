package delivery

import (
	"encoding/json"
	"gateway/internal"
	"gateway/internal/utils"
	"gateway/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type GatewayHandler struct {
	usecase internal.Usecase
}

func NewGatewayHandler(usecase internal.Usecase) *GatewayHandler {
	return &GatewayHandler{usecase: usecase}
}

const layout = "2006-01-02"

func (h *GatewayHandler) GetLibraries(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		utils.Response(w, models.BadRequest, "Invalid type of parameter", nil)
	}
	city := r.URL.Query().Get("city")
	if city == "" {
		utils.Response(w, models.BadRequest, "", nil)
	}

	libs, status := h.usecase.GetLibList(int64(page), int64(size), city)
	body, err := json.Marshal(libs[0])
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
		return
	}
	utils.Response(w, status, "", body)
}

func (h *GatewayHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
		return
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
		return
	}

	vars := mux.Vars(r)
	uidStr := vars["libraryUid"]
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
		return
	}

	var showAll bool
	showAllStr := r.URL.Query().Get("showAll")
	if showAllStr == "true" {
		showAll = true
	} else {
		showAll = false
	}

	books, status := h.usecase.GetBookList(int64(page), int64(size), showAll, uid)
	body, err := json.Marshal(books[0])
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
		return
	}
	utils.Response(w, status, "", body)
}

func (h *GatewayHandler) GetReservations(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	list, st := h.usecase.GetReservationInfo(name)
	if st != models.OK {
		utils.Response(w, st, "", nil)
		return
	}
	listParsed := []models.BookReservationResponseParsed{}
	for _, el := range list {
		node := models.BookReservationResponseParsed{
			ReservationUid: el.ReservationUid,
			Status:         el.Status,
			StartDate:      el.StartDate.Format(layout),
			TillDate:       el.TillDate.Format(layout),
			Book:           el.Book,
			Lib:            el.Lib,
		}
		listParsed = append(listParsed, node)
	}
	body, err := json.Marshal(listParsed)
	if err != nil {
		utils.Response(w, st, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}

func (h *GatewayHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	req := models.TakeBookRequest{}
	reqParsed := models.TakeBookRequestPreParsed{}

	err := json.NewDecoder(r.Body).Decode(&reqParsed)
	if err != nil {
		errorStr := models.ValidationErrorResponse{
			Message: "Incorrect body attached",
		}
		errBody, _ := json.Marshal(errorStr)
		utils.Response(w, models.BadRequest, "", errBody)
		return
	}

	req.BookUid = reqParsed.BookUid
	req.LibraryUid = reqParsed.LibraryUid
	req.TillDate, err = time.Parse(layout, reqParsed.TillDate)

	if err != nil {
		errorStr := models.ValidationErrorResponse{
			Message: "Incorrect body attached",
			Errors: map[string]string{
				"date": "invalid date note",
			},
		}
		errBody, _ := json.Marshal(errorStr)
		utils.Response(w, models.BadRequest, "", errBody)
		return
	}
	validErr := utils.Validate(req)
	if validErr != nil {
		errBody, err := json.Marshal(*validErr)
		if err != nil {
			utils.Response(w, models.InternalError, "", nil)
		}
		utils.Response(w, models.BadRequest, "", errBody)
		return
	}

	book, st := h.usecase.TakeBook(name, req)
	book.StartDate = time.Now()
	if st != models.OK {
		utils.Response(w, st, "", nil)
		return
	}
	bookParsed := models.TakeBookResponseParsedTime{
		ReservationUid: book.ReservationUid,
		Status:         book.Status,
		StartDate:      book.StartDate.Format(layout),
		TillDate:       book.TillDate.Format(layout),
		Book:           book.Book,
		Library:        book.Library,
		Rating:         book.Rating,
	}
	body, err := json.Marshal(bookParsed)
	if err != nil {
		utils.Response(w, st, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}

func (h *GatewayHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	bookReqParsed := models.ReturnBookRequestPreParsed{}
	bookReq := models.ReturnBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&bookReqParsed)

	bookReq.Condition = bookReqParsed.Condition
	bookReq.Date, err = time.Parse(layout, bookReqParsed.Date)

	name := r.Header.Get("X-User-Name")
	vars := mux.Vars(r)
	UidStr := vars["reservationUid"]
	Uid, err := uuid.Parse(UidStr)
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
		return
	}
	status := h.usecase.ReturnBook(Uid, name, bookReq)
	if status == models.NotFound {
		errBody, err := json.Marshal(models.ErrorResponse{Message: "Not Found"})
		if err != nil {
			utils.Response(w, models.InternalError, "", nil)
		}
		utils.Response(w, status, "", errBody)
		return
	}
	utils.Response(w, status, "", nil)
}

func (h *GatewayHandler) GetRating(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	resp, st := h.usecase.GetRating(name)
	if st != models.OK {
		utils.Response(w, models.InternalError, "", nil)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}
