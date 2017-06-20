package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			})
	})
	initializeRoutes(router)
	router.Run(":3000")
}
