package config

import "os"

type Config struct {
	HTTPAddr string
}

func Read() Config {
	var c Config
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		c.HTTPAddr = addr
	}
	return c
}
