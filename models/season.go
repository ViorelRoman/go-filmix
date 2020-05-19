package models

import (
	"time"
)

type Season struct {
	Title       string
	Description string
	DetailUrl   string
	ImageUrl    string
	PubDate     *time.Time
}

func (s *Season) GetEpisodes() ([]Episode, error) {
	return []Episode{}, nil
}
