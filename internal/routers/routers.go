package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter() http.Handler {

	r := gin.Default()

	// Favicon
	r.StaticFile("/favicon.ico", "assets/favicon.ico")

	// Set the directory to load html templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/ping", pingHandler)
	r.GET("/health", healthHandler)

	return r
}
