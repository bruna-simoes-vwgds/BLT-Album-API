package server

import (
	"album-app/api/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	loadEndpoints(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error running gin server: ", err)
		return
	}
}

func loadEndpoints(router *gin.Engine) {
	routers.LoadAlbumRoutes(router)
}
