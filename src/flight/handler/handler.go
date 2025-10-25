package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	services "github.com/Mamvriyskiy/lab2-template/src/flight/services"
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

	flight := router.Group("")

	flight.GET("/flight", h.GetInfoAboutFlight)

	return router
}