package server

import (
	"album-app/cmd/api/handlers"
	"album-app/cmd/api/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(handlers.SetContentType())
	router.NoRoute(handlers.NotFoundError())

	loadEndpoints(router)

	err := router.Run(":" + os.Getenv("GO_LISTENING_PORT"))
	if err != nil {
		log.Fatal("Error running gin server: ", err)
		return
	}
}

func loadEndpoints(router *gin.Engine) {
	routers.LoadAlbumRoutes(router)
}
