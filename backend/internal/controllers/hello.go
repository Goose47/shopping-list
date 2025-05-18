package controllers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type HelloController struct {
	log *slog.Logger
}

func NewHelloController(
	log *slog.Logger,
) *HelloController {
	return &HelloController{
		log: log,
	}
}

func (con *HelloController) Hello(c *gin.Context) {
	res := "hello, world"
	c.JSON(http.StatusOK, res)
}
