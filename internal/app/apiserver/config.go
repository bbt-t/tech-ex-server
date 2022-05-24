package apiserver

import "tech-ex-server/internal/app/storage"

type Config struct {
	BinAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}
}
