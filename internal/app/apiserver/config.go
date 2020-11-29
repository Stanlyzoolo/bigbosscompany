package apiserver

import "github.com/Stanlyzoolo/bigbosscompany/store"

// Config...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

// NewConfig...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Store:    store.NewConfig(),
	}
}
