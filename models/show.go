package models

type Show struct {
	title       string
	description string
	DetailUrl   string
	Show        *Show
}

func (s *Show) GetSeasons() ([]Season, error) {
	return []Season{}, nil
}
