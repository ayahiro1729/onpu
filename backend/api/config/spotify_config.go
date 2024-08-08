package config

import "github.com/caarlos0/env"

type SpotifyConfig struct {
	ClientID     string `env:"SPOTIFY_CLIENT_ID"`
	ClientSecret string `env:"SPOTIFY_CLIENT_SECRET"`
	RedirectURI  string `env:"SPOTIFY_REDIRECT_URI"`
	JWTSecret    string `env:"JWT_SECRET"`
}

func NewSpotifyConfig() (*SpotifyConfig, error) {
	cfg := &SpotifyConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
