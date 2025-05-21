package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"shopping-list/internal/lib/auth"
)

type Auth struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewAuth(
	log *slog.Logger,
	db *gorm.DB,
) *Auth {
	return &Auth{
		log: log,
		db:  db,
	}
}

func (con *Auth) Me(c *gin.Context) {
	const op = "AuthController.Me"
	log := con.log.With(slog.String("op", op))

	user, err := auth.User(c)
	if err != nil {
		log.Error("failed to retrieve user", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
