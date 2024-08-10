package types

// Spotify API User's Top Tracks
type TrackItem struct {
	Name   string `json:"name"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	ExternalURL struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}
