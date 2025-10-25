package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func (h *Handler) GetInfoAboutUserPrivilege(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
    if username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
        return
    }

    resp, err := h.services.GetInfoAboutUserPrivilege(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


func (h *Handler) UpdateBonus(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
    if username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header is required"})
        return
    }

    ticketUid := c.Param("ticketUid")
    if ticketUid == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ticketUid is required"})
        return
    }

    err := h.services.UpdateBonus(username, ticketUid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}
