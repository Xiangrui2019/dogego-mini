package routers

import (
	"dogego-mini/api"
	"dogego-mini/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 中间件, 顺序必须是这样
	//router.Use(middlewares.SentryReportor())
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)
	}

	return router
}
