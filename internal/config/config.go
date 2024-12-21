package config

import "os"

var (
	service  = os.Getenv("SERVICE")
	version  = os.Getenv("VERSION")
	color    = os.Getenv("COLOR")
	webPort1 = func() string {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		return port
	}()
	webPort2 = func() string {
		port := os.Getenv("PORT_2")
		if port == "" {
			port = "3000"
		}
		return port
	}()
)

type Config struct {
	Color   string `json:"color,omitempty"`
	Port    string `json:"port,omitempty"`
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
}

func WebConfig1() *Config {
	cfg := &Config{
		Color:   color,
		Port:    webPort1,
		Service: service,
		Version: version,
	}
	return cfg
}

func WebConfig2() *Config {
	cfg := WebConfig1()
	cfg.Port = webPort2
	return cfg
}
