package server

import (
	"album-app/internal/api/database"
	"album-app/internal/api/handlers"
)

type App struct {
	AlbumHandler *handlers.AlbumHandler
}

func buildApp() *App {
	dbHandler := database.Handler{}
	dbHandler.InitializeMongoDB()

	albumHandler := handlers.NewAlbumHandler(&dbHandler)

	return &App{
		AlbumHandler: albumHandler,
	}
}
