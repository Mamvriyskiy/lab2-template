package handler

import (
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

	bonus := router.Group("")

	bonus.GET("/flight", h.GetInfoAboutFlight)


	return router
}