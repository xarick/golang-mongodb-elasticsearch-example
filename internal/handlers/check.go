package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckHandler struct{}

func NewCheckHandler() *CheckHandler {
	return &CheckHandler{}
}

func (h *CheckHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
