package config

import "os"

const (
	DEFAULT_PRIMARY_WEB_PORT   = "8080"
	DEFAULT_SECONDARY_WEB_PORT = "3000"
	DEFAULT_PRIMARY_GRPC_PORT  = "9090"

	PRIMARY_WEB_PORT_ENV   = "PORT"
	SECONDARY_WEB_PORT_ENV = "PORT_2"
	PRIMARY_GRPC_PORT_ENV  = "GRPC_PORT"
)

var (
	service  = os.Getenv("SERVICE")
	version  = os.Getenv("VERSION")
	color    = os.Getenv("COLOR")
	webPort1 = func() string {
		port := os.Getenv(PRIMARY_WEB_PORT_ENV)
		if port == "" {
			port = DEFAULT_PRIMARY_WEB_PORT
		}
		return port
	}()
	webPort2 = func() string {
		port := os.Getenv(SECONDARY_WEB_PORT_ENV)
		if port == "" {
			port = DEFAULT_SECONDARY_WEB_PORT
		}
		return port
	}()
	grpcPort1 = func() string {
		port := os.Getenv(PRIMARY_GRPC_PORT_ENV)
		if port == "" {
			port = DEFAULT_PRIMARY_GRPC_PORT
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

func GrpcConfig1() *Config {
	cfg := WebConfig1()
	cfg.Port = grpcPort1
	return cfg
}
