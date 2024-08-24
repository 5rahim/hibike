package tvdb

func (s *ExtendedSeason) IsAbsolute() bool {
	return s.Type.Type == "absolute" && s.Number == 1
}

func (s *ExtendedSeason) IsSpecialsAndMovies() bool {
	return s.Type.Type == "official" && s.Number == 0
}
