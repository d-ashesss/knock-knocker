package main

import (
	"github.com/joeshaw/envdecode"
	"log"
)

type Config struct {
	Port string `env:"PORT,default=8080"`
}

func NewConfig() *Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}
