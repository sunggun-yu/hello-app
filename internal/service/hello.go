package hello

import (
	"fmt"
	"os"
	"time"

	"github.com/sunggun-yu/hello-app/internal/config"
	"github.com/sunggun-yu/hello-app/internal/models"
)

var (
	name        = os.Getenv("NAME")
	hostname, _ = os.Hostname()
)

const (
	healthy = "HEALTHY"
)

func Index(config *config.Config) models.Hello {
	return models.Hello{
		Color:     config.Color,
		Service:   config.Service,
		Version:   config.Version,
		Instance:  hostname,
		Port:      config.Port,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func Hello() string {
	message := "Hello!"
	if name != "" {
		message = fmt.Sprintf("Hello, %s!", name)
	}
	if hostname != "" {
		message = fmt.Sprintf("%s, Instance: %s", message, hostname)
	}
	return message
}

func Ping(config *config.Config) models.Hello {
	return models.Hello{
		Message:   "pong",
		Service:   config.Service,
		Version:   config.Version,
		Instance:  hostname,
		Port:      config.Port,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func Health() string {
	return healthy
}
