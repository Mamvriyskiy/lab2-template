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
    targetURL := "http://localhost:8070/ticket/"

    ticketUid := c.Param("X-User-Name")
    if ticketUid == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
        return
    }

    req, err := http.NewRequest("GET", targetURL, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    targetURL += ticketUid

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}


func (h *Handler) GetInfoAboutAllUserTickets(c *gin.Context) {
    targetURL := "http://localhost:8070/tickets"

    username := c.GetHeader("X-User-Name")
    if username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
        return
    }

    req, err := http.NewRequest("GET", targetURL, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    req.Header.Set("X-User-Name", username)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}


func (h *Handler) GetInfoAboutUserPrivilege(c *gin.Context) {
    targetURL := "http://localhost:8050/privilege"

    username := c.GetHeader("X-User-Name")
    if username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
        return
    }

    req, err := http.NewRequest("GET", targetURL, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    req.Header.Set("X-User-Name", username)

    client := &http.Client{}
    resp, err := client.Do(req)
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

