package routers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunggun-yu/hello-app/internal/config"
	helloService "github.com/sunggun-yu/hello-app/internal/service"
)

var (
	port = config.Config.Port
)

// indexHandler
func indexHandler(c *gin.Context) {
	data := helloService.Index()
	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
		"color":     data.Color,
		"service":   data.Service,
		"version":   data.Version,
		"instance":  data.Instance,
		"host":      c.Request.Host,
		"port":      port,
		"timestamp": data.Timestamp,
	})
}

// helloHandler
func helloHandler(c *gin.Context) {

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

	message := helloService.Hello()
	if duration > 0 {
		message = fmt.Sprintf("%s, waited %v\n", message, duration)
	}
	c.String(http.StatusOK, message)
}

// pingHandler
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, helloService.Ping())
}

// healthHandler
func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, helloService.Health())
}
