package server

import (
	"album-app/internal/api/models"
	"album-app/internal/api/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(SetContentType())
	router.NoRoute(NotFoundError())

	app := buildApp()
	loadEndpoints(router, app)

	err := router.Run(":" + os.Getenv("GO_LISTENING_PORT"))
	if err != nil {
		log.Fatal("Error running gin server: ", err)
		return
	}
}

func SetContentType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}

func NotFoundError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Error{Message: "Method not found"})
	}
}

func loadEndpoints(router *gin.Engine, app *App) {
	v1 := router.Group("/api/v1")
	{
		routes.LoadAlbumRoutes(v1, *app.AlbumHandler)
	}
}
