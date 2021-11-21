package server

import (
	"net/http"
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
			res, err := handle()
			// fmt.Printf("%+v\n", res)
			if err != nil {
				ErrorResponse(c, err)
			} else {
				OkResponse(c, res)
			}

		})

		group.POST("/set", func(c *gin.Context) {
			c.JSON(http.StatusOK, "set")
		})
	}

	return router
}

type Key struct {
	Name    string "json:name"
	Surname string "json:surname"
}

func handle() (*Key, error) {
	// op := "Handle"
	//err := &AppError{Code: constants.EINVALID, Op: op, Message: "asdasdasd"}
	return &Key{Name: "fatih", Surname: "kurt"}, nil
}
