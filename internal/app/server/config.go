package server

import "github.com/Heroin-lab/diplom.git/internal/app/database"

type Config struct {
	BindAddr       string `toml:"bind_addr"`
	LogLevel       string `toml:"log_level"`
	DatabaseConfig *database.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:       ":7777",
		LogLevel:       "debug",
		DatabaseConfig: database.NewConfig(),
	}
}
