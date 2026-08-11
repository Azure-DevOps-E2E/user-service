package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"polyglot-shop/user-service/internal/requestid"
	"polyglot-shop/user-service/internal/user"
)

func New() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), requestid.Middleware())

	repository := user.NewRepository()
	handler := user.NewHandler(repository)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP", "service": "user-service"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", handler.List)
		v1.GET("/users/:id", handler.Get)
	}

	return router
}
