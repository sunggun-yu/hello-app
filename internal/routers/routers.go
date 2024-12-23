package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunggun-yu/hello-app/internal/config"
)

func DefaultRouter(config *config.Config) http.Handler {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(customLoggerMiddleware(config))
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Printf("[GIN-debug][web:%v] %v\t%v\t --> %v (%v handlers)\n", config.Port, httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// Favicon
	r.StaticFile("/favicon.ico", "assets/favicon.ico")

	// Set the directory to load html templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexHandler(config))
	r.GET("/hello", helloHandler)
	r.GET("/ping", pingHandler(config))
	r.GET("/health", healthHandler)

	return r
}

func customLoggerMiddleware(config *config.Config) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Default Gin format:
		// [GIN] 31/Mar/2024 - 12:34:56 | 200 | 123.456µs | ::1 | GET "/path"

		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("%v [web:%v]\t|%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			config.Port,
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	})
}
