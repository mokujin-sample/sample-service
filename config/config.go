package config

import (
	"service/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Environment string
	Port        string
	Database    *Database
	DatabaseCH  *Database
	RabbitMQ    *RabbitMQ
}

// Database is a struct that contains DB's configuration variables
type Database struct {
	Host     string
	User     string
	DB       string
	Password string
	Port     string
}

type RabbitMQ struct {
	DSN         string
	NotifyQueue string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	port := env.MustGet("PORT")
	// set default PORT if missing
	if port == "" {
		port = "3000"
	}
	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        port,
		Database: &Database{
			Host:     env.MustGet("DATABASE_HOST"),
			User:     env.MustGet("DATABASE_USER"),
			DB:       env.MustGet("DATABASE_DB"),
			Password: env.MustGet("DATABASE_PASSWORD"),
		},
		RabbitMQ: &RabbitMQ{
			DSN:         env.MustGet("RABBIT_DSN"),
			NotifyQueue: env.MustGet("RABBIT_NOTIFY_QUEUE"),
		},
		DatabaseCH: &Database{
			Host: env.MustGet("DATABASE_CH_HOST"),
			DB:   env.MustGet("DATABASE_CH_DB"),
			Port: env.MustGet("DATABASE_CH_PORT"),
		},
	}, nil
}
