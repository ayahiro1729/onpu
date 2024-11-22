package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type SpotifyConfig struct {
	ClientID     string `env:"SPOTIFY_CLIENT_ID"`
	ClientSecret string `env:"SPOTIFY_CLIENT_SECRET"`
	RedirectURI  string `env:"SPOTIFY_REDIRECT_URI"`
}

func NewSpotifyConfig() (*SpotifyConfig, error) {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &SpotifyConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
