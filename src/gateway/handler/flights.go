package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetInfoAboutFlight(c *gin.Context) {
    targetURL := "http://localhost:8060/flight"

    resp, err := http.Get(targetURL)
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}


func (h *Handler) GetInfoAboutUser(c *gin.Context) {

}


func (h *Handler) GetInfoAboutUserTicket(c *gin.Context) {

}


func (h *Handler) GetInfoAboutAllUserTickets(c *gin.Context) {

}


func (h *Handler) GetInfoAboutUserPrivilege(c *gin.Context) {
	targetURL := "http://localhost:8050/privilege"

    resp, err := http.Get(targetURL)
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}


func (h *Handler) BuyTicketUSer(c *gin.Context) {

}


func (h *Handler) DeleteTicketUSer(c *gin.Context) {

}

