package config

import "github.com/martin-harano/kong-vs-krakend/mockapp/pkg/env"

type Config struct {
	Server struct {
		Port     string
		UseSonic bool
		Prefork  bool
	}

	Routers struct {
		Path string
	}
}

func NewConfig() *Config {
	return &Config{
		Server: struct {
			Port     string
			UseSonic bool
			Prefork  bool
		}{
			Port:     env.GetEnvOrDie("SERVER_PORT"),
			UseSonic: env.GetEnvOrDie("ENABLE_SONIC_JSON") == "1",
			Prefork:  env.GetEnvOrDie("ENABLE_PREFORK") == "1",
		},
		Routers: struct {
			Path string
		}{
			Path: env.GetEnvOrDie("PATH_NAME"),
		},
	}
}
