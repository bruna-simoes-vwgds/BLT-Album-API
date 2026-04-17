package server

import (
	"album-app/internal/api/handlers"
	"album-app/internal/api/models"
	"album-app/internal/api/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func Run() {
	loadEnvVariablesFromEnvFile()

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
		ctx.JSON(http.StatusNotFound, models.Error{Message: "Method not found", Timestamp: time.Now().String()})
	}
}

func loadEndpoints(router *gin.Engine, app *App) {
	v1 := router.Group("/api/v1")
	{
		routes.LoadAlbumRoutes(v1, *app.AlbumHandler)
		routes.HealthCheck(v1, handlers.GeneralHandler{})
	}
}

// This function loads environment variables from a .env file, but only if the application is not running inside a Docker container.
// If the application is running inside a Docker container, it assumes that the environment variables are already set and does not attempt to load them from a file.
func loadEnvVariablesFromEnvFile() {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
