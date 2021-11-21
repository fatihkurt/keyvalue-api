package server

import (
	"deliveryhero/handler"
	"deliveryhero/helper"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Server running.")
	})

	group := router.Group("/api")
	{
		group.GET("/get", func(c *gin.Context) {
			h := handler.NewKeyValueHandler(c)
			err := h.HanldeGetKey()
			if err != nil {
				helper.ErrorResponse(c, err)
			}
		})

		group.POST("/set", func(c *gin.Context) {
			h := handler.NewKeyValueHandler(c)
			err := h.HanldeSetKey()
			if err != nil {
				helper.ErrorResponse(c, err)
			}
		})
	}

	return router
}
