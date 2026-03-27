package handlers

import (
	"album-app/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	id := c.Param("id")
	album, foundAlbum := service.GetAlbumById(id)

	if foundAlbum {
		c.JSON(http.StatusOK, album)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}
