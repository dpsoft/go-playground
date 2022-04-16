package settings

import "flag"

type Settings struct {
	Enabled bool
	Server  Server
}

type Server struct {
	Port string
	Host string
}

func NewSettings() Settings {
	var enabled bool
	var port string
	var host string

	flag.BoolVar(&enabled, "enabled", true, "Enable the server")
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.StringVar(&host, "host", "localhost", "Host to listen on")

	flag.Parse()

	return Settings{
		Enabled: enabled,
		Server:  Server{Port: port, Host: host},
	}
}
