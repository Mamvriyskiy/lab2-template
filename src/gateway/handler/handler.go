package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	services "github.com/Mamvriyskiy/lab2-template/src/gateway/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.GET("/manage/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	flight := router.Group("api/v1/")

	// Получить список всех перелетов
	flight.GET("/flight", h.GetInfoAboutFlight)

	// Возвращается информация о билетах и статусе в системе привилегии
	flight.GET("/me", h.GetInfoAboutUser)

	// Получить информацию о всех билетах пользователя
	flight.GET("/tickets", h.GetInfoAboutAllUserTickets)

	// Получить информацию о конкретном билете пользователя
	flight.GET("/ticket/:ticketUid", h.GetInfoAboutUserTicket)

	// Покупка билета
	flight.POST("/tickets/:ticketUid", h.BuyTicketUSer)

	// Возврат билета
	flight.DELETE("/tickets/:ticketUid", h.DeleteTicketUSer)

	// Получить информацию о состоянии бонусного счета
	flight.GET("/privilege", h.GetInfoAboutUserPrivilege)

	return router
}