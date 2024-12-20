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

func Index() models.Hello {
	return models.Hello{
		Color:     config.Config.Color,
		Service:   config.Config.Service,
		Version:   config.Config.Version,
		Instance:  hostname,
		Port:      config.Config.Port,
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

func Ping() models.Hello {
	return models.Hello{
		Message:   "pong",
		Service:   config.Config.Service,
		Version:   config.Config.Version,
		Instance:  hostname,
		Port:      config.Config.Port,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func Health() string {
	return healthy
}
