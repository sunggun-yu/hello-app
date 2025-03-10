package models

type Hello struct {
	Color      string `json:"color,omitempty"`
	Host       string `json:"host,omitempty"`
	Instance   string `json:"instance,omitempty"`
	Message    string `json:"message,omitempty"`
	Name       string `json:"name,omitempty"`
	Port       string `json:"port,omitempty"`
	Service    string `json:"service,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	Version    string `json:"version,omitempty"`
	ClientIP   string `json:"clientIp,omitempty"`
	RemoteAddr string `json:"remoteAddr,omitempty"`
}
