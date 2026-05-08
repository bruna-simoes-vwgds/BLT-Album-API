package server

import (
	"album-app/internal/api/database"
	"album-app/internal/api/handlers"
)

type App struct {
	DatabaseHandler *database.DBHandler
	AlbumHandler    *handlers.AlbumHandler
	SongHandler     *handlers.SongHandler
}

func buildApp() *App {
	dbHandler := database.DBHandler{}
	dbHandler.InitializeMongoDB()

	albumHandler := handlers.NewAlbumHandler(&dbHandler)
	songHandler := handlers.NewSongHandler(&dbHandler)

	return &App{
		DatabaseHandler: &dbHandler,
		AlbumHandler:    albumHandler,
		SongHandler:     songHandler,
	}
}
