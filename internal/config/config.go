package config

import "os"

var (
	service = os.Getenv("SERVICE")
	version = os.Getenv("VERSION")
	color   = os.Getenv("COLOR")
	port    = func() string {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		return port
	}()

	Config *config
)

type config struct {
	Color   string `json:"color,omitempty"`
	Port    string `json:"port,omitempty"`
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
}

func init() {
	Config = &config{
		Color:   color,
		Port:    port,
		Service: service,
		Version: version,
	}
}
