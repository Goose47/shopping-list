// Package app defines application model.
package app

import (
	"context"
	"log/slog"
	serverapp "shopping-list/internal/app/server"
	"shopping-list/internal/background"
	"shopping-list/internal/config"
	"shopping-list/internal/controllers"

	//"shopping-list/internal/db/mysql"
	"shopping-list/internal/db/postgres"
	"shopping-list/internal/server"
)

type clearer interface {
	Clear(ctx context.Context)
}

// App represents application.
type App struct {
	Server      *serverapp.Server
	TaskManager *background.TaskManager
	Cache       clearer
}

// New creates all dependencies for App and returns new App instance.
func New(
	log *slog.Logger,
	env string,
	port int,
	dbConfig config.DB,
) *App {
	//db, err := mysql.New(dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Pass, dbConfig.DBName)
	db, err := postgres.New(dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Pass, dbConfig.DBName)
	if err != nil {
		panic(err)
	}

	helloCon := controllers.NewHelloController(log)
	authCon := controllers.NewAuth(log, db)

	router := server.NewRouter(
		env,
		helloCon,
		authCon,
	)
	serverApp := serverapp.New(log, port, router)

	tm := background.NewTaskManager(log, db)

	return &App{
		Server:      serverApp,
		TaskManager: tm,
	}
}
