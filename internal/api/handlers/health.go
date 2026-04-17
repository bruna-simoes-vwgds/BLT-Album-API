package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralHandler struct {
}

func (generalHandler *GeneralHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "healthy"})
}
