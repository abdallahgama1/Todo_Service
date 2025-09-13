package config

import "errors"

type Config struct {
	DSN string
}

func LoadConfig() (*Config, error) {

	dsn := "postgres://your_user:your_password@localhost:5432/"
	if dsn == "" {
		return nil, errors.New("database DSN not configured")
	}

	return &Config{DSN: dsn}, nil
}
