package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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
			"color":     color,
			"service":   serviceName,
			"version":   version,
			"instance":  hostname,
			"host":      c.Request.Host,
			"timestamp": time.Now().Format(time.RFC3339),
		}
		c.HTML(http.StatusOK, "index.html.tmpl", data)
	})

	r.GET("/hello", func(c *gin.Context) {

		// Get the 'wait' parameter from query string
		waitStr := c.Query("wait")

		var duration time.Duration

		if waitStr != "" {
			// Try to parse as duration
			d, err := time.ParseDuration(waitStr)
			if err != nil {
				// If parsing fails, try to parse as integer and assume seconds
				sec, err := strconv.Atoi(waitStr)
				if err != nil {
					// Unable to parse, set duration to 0
					duration = 0
				} else {
					duration = time.Duration(sec) * time.Second
				}
			} else {
				duration = d
			}
		}

		if duration > 0 {
			time.Sleep(duration)
		}

		message := "Hello!"
		if name != "" {
			message = fmt.Sprintf("Hello, %s!", name)
		}
		if duration > 0 {
			message = fmt.Sprintf("%s, waited %v, Instance: %s\n", message, duration, hostname)
		} else {
			message = fmt.Sprintf("%s, Instance: %s\n", message, hostname)
		}
		c.String(http.StatusOK, message)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"service":   serviceName,
			"version":   version,
			"instance":  hostname,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "HEALTHY")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
