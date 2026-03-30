package routers

import (
	"album-app/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAlbumRoutes(router *gin.Engine) {
	router.GET("/albums/id/:id", handlers.GetAlbums)
}
