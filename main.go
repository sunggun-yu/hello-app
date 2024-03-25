package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	serviceName := os.Getenv("MY_NAME")
	r.GET("/", func(c *gin.Context) {
		if serviceName != "" {
			c.String(http.StatusOK, "hello world from "+serviceName)
		} else {
			c.String(http.StatusOK, "hello world")
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
