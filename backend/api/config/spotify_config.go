package config

import "github.com/caarlos0/env"

type SpotifyConfig struct {
  ClientID     string `env:"SPOTIFY_CLIENT_ID"`
  ClientSecret string `env:"SPOTIFY_CLIENT_SECRET"`
  RedirectURI  string `env:"SPOTIFY_REDIRECT_URI"`
}

func NewSpotifyConfig() *SpotifyConfig {
	cfg := &SpotifyConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
