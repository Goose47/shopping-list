// application entrypoint
package main

import (
	"log/slog"
	apppkg "shopping-list/internal/app"
	"shopping-list/internal/config"
	"shopping-list/internal/logger"
	telegramapp "shopping-list/internal/telegram/app"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	app := apppkg.New(log, cfg.Env, cfg.Port, cfg.DB, cfg.Telegram)

	app.TaskManager.Run()

	telegram := telegramapp.New(log, cfg.Telegram.Token, cfg.AppURL)
	go func() {
		err := telegram.Run()
		log.Error("failed to run telegram", slog.Any("error", err))
	}()

	err := app.Server.Serve()
	log.Error("application has stopped", slog.Any("error", err))
}
