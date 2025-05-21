// Package server defines router settings and application endpoints.
package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"shopping-list/internal/controllers"
	envpkg "shopping-list/internal/domain/enums/env"
	"shopping-list/internal/server/middleware"
)

// NewRouter sets router mode based on env, registers middleware, defines handlers and options and creates new gin router.
func NewRouter(
	log *slog.Logger,
	env string,
	db *gorm.DB,
	botSecret string,
	helloCon *controllers.HelloController,
	authCon *controllers.Auth,
) *gin.Engine {
	var mode string
	switch env {
	case envpkg.Local:
	case envpkg.Dev:
		mode = gin.DebugMode
	case envpkg.Prod:
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	r := gin.New()

	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	r.Use(gin.Recovery())

	r.GET("hello", helloCon.Hello)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		protected := v1.Group("")
		protected.Use(middleware.TelegramAuthMiddleware(log, db, botSecret))

		auth := protected.Group("/auth")

		auth.GET("me", authCon.Me)
	}

	return r
}
