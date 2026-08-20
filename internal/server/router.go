package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"nexuscart/user-service/internal/requestid"
	"nexuscart/user-service/internal/user"
)

const defaultServiceVersion = "1.0.0"

func serviceVersion() string {
	if version := strings.TrimSpace(os.Getenv("APP_VERSION")); version != "" {
		return version
	}
	return defaultServiceVersion
}

func serviceImageTag() string {
	if imageTag := strings.TrimSpace(os.Getenv("APP_IMAGE_TAG")); imageTag != "" {
		return imageTag
	}
	return serviceVersion()
}

func New() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), requestid.Middleware())

	repository := user.NewRepository()
	handler := user.NewHandler(repository)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "UP",
			"service":  "user-service",
			"version":  serviceVersion(),
			"imageTag": serviceImageTag(),
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", handler.List)
		v1.GET("/users/:id", handler.Get)
	}

	return router
}
