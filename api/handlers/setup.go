package handlers

import (
	"album-app/api/datamodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetContentType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}

func NotFoundError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, datamodel.Error{Message: "Method not found", Code: 404})
	}
}
