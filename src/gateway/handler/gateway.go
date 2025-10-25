package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	modelGateway "github.com/Mamvriyskiy/lab2-template/src/gateway/model"
	"github.com/gin-gonic/gin"
)

func forwardRequest(c *gin.Context, method, targetURL string, headers map[string]string) (int, []byte, http.Header, error) {
	if len(c.Request.URL.RawQuery) > 0 {
		targetURL = fmt.Sprintf("%s?%s", targetURL, c.Request.URL.RawQuery)
	}

	req, err := http.NewRequest(method, targetURL, nil)
	if err != nil {
		return 0, nil, nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, resp.Header, err
	}

	return resp.StatusCode, body, resp.Header, nil
}

func (h *Handler) GetInfoAboutFlight(c *gin.Context) {
	status, body, headers, err := forwardRequest(c, "GET", "http://flight:8060/flight", nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.Data(status, headers.Get("Content-Type"), body)
}

func (h *Handler) GetInfoAboutUserTicket(c *gin.Context) {
	ticketUid := c.Param("ticketUid")
	if ticketUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticketUid is required"})
		return
	}

	// 1️⃣ Запрашиваем билет
	ticketURL := "http://ticket:8070/ticket/" + ticketUid
	status, body, _, err := forwardRequest(c, "GET", ticketURL, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if status != http.StatusOK {
		c.Data(status, "application/json", body)
		return
	}

	var ticket modelGateway.Ticket
	if err := json.Unmarshal(body, &ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse ticket response"})
		return
	}

	// 2️⃣ Запрашиваем данные о рейсе
	flightURL := "http://flight:8060/flight/" + ticket.FlightNumber
	flightStatus, flightBody, _, err := forwardRequest(c, "GET", flightURL, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if flightStatus != http.StatusOK {
		c.Data(flightStatus, "application/json", flightBody)
		return
	}

	var flight modelGateway.Flight
	if err := json.Unmarshal(flightBody, &flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse flight response"})
		return
	}

	// 3️⃣ Объединяем оба ответа
	response := modelGateway.TicketInfo{
		TicketUID:    ticket.TicketUID,
		FlightNumber: flight.FlightNumber,
		FromAirport:  flight.FromAirport,
		ToAirport:    flight.ToAirport,
		Date:         flight.Datetime,
		Price:        flight.Price,
		Status:       ticket.Status,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetInfoAboutAllUserTickets(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
		return
	}

	// 1️⃣ Получаем все билеты пользователя
	headers := map[string]string{"X-User-Name": username}
	status, body, respHeaders, err := forwardRequest(c, "GET", "http://localhost:8070/tickets", headers)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if status != http.StatusOK {
		c.Data(status, respHeaders.Get("Content-Type"), body)
		return
	}

	var tickets []modelGateway.Ticket
	if err := json.Unmarshal(body, &tickets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse tickets"})
		return
	}

	print(tickets)
	// 2️⃣ Проходим по каждому билету и запрашиваем информацию о рейсе
	var ticketInfos []modelGateway.TicketInfo
	for _, ticket := range tickets {
		if ticket.FlightNumber == "" {
			continue
		}
		flightURL := "http://localhost:8060/flight/" + ticket.FlightNumber
		flightStatus, flightBody, _, err := forwardRequest(c, "GET", flightURL, nil)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if flightStatus != http.StatusOK {
			c.Data(flightStatus, "application/json", flightBody)
			return
		}

		var flight modelGateway.Flight
		if err := json.Unmarshal(flightBody, &flight); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse flight response"})
			return
		}

		print(flight.FlightNumber)
		// 3️⃣ Объединяем данные билета и рейса
		ticketInfos = append(ticketInfos, modelGateway.TicketInfo{
			TicketUID:    ticket.TicketUID,
			FlightNumber: flight.FlightNumber,
			FromAirport:  flight.FromAirport,
			ToAirport:    flight.ToAirport,
			Date:         flight.Datetime,
			Price:        flight.Price,
			Status:       ticket.Status,
		})
	}

	// 4️⃣ Отправляем массив в ответе
	c.JSON(http.StatusOK, ticketInfos)
}

func (h *Handler) GetInfoAboutUserPrivilege(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
		return
	}

	headers := map[string]string{"X-User-Name": username}
	status, body, respHeaders, err := forwardRequest(c, "GET", "http://bonus:8050/privilege", headers)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.Data(status, respHeaders.Get("Content-Type"), body)
}

func (h *Handler) GetInfoAboutUser(c *gin.Context) {

}

func (h *Handler) BuyTicketUSer(c *gin.Context) {

}

func (h *Handler) DeleteTicketUSer(c *gin.Context) {
	ticketUid := c.Param("ticketUid")
	if ticketUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticketUid is required"})
		return
	}

	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
		return
	}

	// 1️⃣ Запрашиваем билет
	ticketURL := "http://ticket:8070/ticket/" + ticketUid
	status, body, _, err := forwardRequest(c, "PATCH", ticketURL, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if status != http.StatusOK {
		c.Data(status, "application/json", body)
		return
	}

	headers := map[string]string{"X-User-Name": username}
	// 2️⃣ Запрашиваем данные о рейсе
	flightURL := "http://flight:8060/bonus/" + ticketUid
	flightStatus, flightBody, _, err := forwardRequest(c, "PATCH", flightURL, headers)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if flightStatus != http.StatusOK {
		c.Data(flightStatus, "application/json", flightBody)
		return
	}

	var flight modelGateway.Flight
	if err := json.Unmarshal(flightBody, &flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse flight response"})
		return
	}
}
