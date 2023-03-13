package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP
		PSQL
	}

	HTTP struct {
	}

	PSQL struct {
		URL string `env:"PSQL_URL"`
	}
)

func New() (*Config, error) {
	var (
		cfg = new(Config)
		err error
	)

	if err = cleanenv.ReadConfig("./config/config.yaml", cfg); err != nil {
		return nil, err
	}

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
