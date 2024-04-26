package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Set the directory from which to load templates
	r.LoadHTMLGlob("templates/*")

	serviceName := os.Getenv("SERVICE")
	version := os.Getenv("VERSION")
	name := os.Getenv("NAME")
	hostname, _ := os.Hostname()
	color := os.Getenv("COLOR")

	r.GET("/", func(c *gin.Context) {

		data := map[string]interface{}{
			"color":    color,
			"service":  serviceName,
			"version":  version,
			"instance": hostname,
			"host":     c.Request.Host,
		}
		c.HTML(http.StatusOK, "index.html.tmpl", data)
	})

	r.GET("/hello", func(c *gin.Context) {

		message := "Hello!"
		if name != "" {
			message = fmt.Sprintf("Hello, %s!", name)
		}
		message = fmt.Sprintf("%s, Instance: %s\n", message, hostname)
		c.String(http.StatusOK, message)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"instance": hostname,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "HEALTHY")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
