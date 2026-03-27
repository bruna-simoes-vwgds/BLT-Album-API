package server

import (
	"album-app/api/handlers"
	"album-app/api/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(handlers.SetContentType())
	router.NoRoute(handlers.NotFoundError())

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
