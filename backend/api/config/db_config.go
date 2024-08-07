package config

import "github.com/caarlos0/env"

type DBConfig struct {
	Env        string `env:"ENV" envDefault:"dev"`
	Port       string `env:"PORT" envDefault:"5432"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBUser     string `env:"DB_USER" envDefault:"user"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DBName     string `env:"DB_NAME" envDefault:"onpu"`
}

func NewDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
