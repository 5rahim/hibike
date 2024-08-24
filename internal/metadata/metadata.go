package metadata

import (
	"fmt"
	"github.com/5rahim/hibike/pkg/animelists"
	"github.com/rs/zerolog"
)

type (
	SeriesMetadata struct {
		Titles    map[string]string           `json:"titles"`
		Episodes  map[string]*EpisodeMetadata `json:"episodes"`
		AnilistID int                         `json:"anilistID"`
		TvdbID    int                         `json:"tvdbID"`
		AnidbID   int                         `json:"anidbID"`
		// Based on the AniDB entry.
		EpisodeCount int `json:"episodeCount"`
		// Based on the AniDB entry.
		SpecialCount int `json:"specialCount"`
		// Based on the AniDB entry.
		SeasonNumber int `json:"seasonNumber"`
		// AbsoluteOffset is the total number of episodes from previous seasons.
		// e.g. if season 1 has 12 episodes, season 2 has 12 episodes and this is the 3rd season, AbsoluteOffset = 24
		AbsoluteOffset int `json:"absoluteOffset"`
	}

	EpisodeMetadata struct {
		// Season number based on the AniDB entry.
		SeasonNumber int `json:"seasonNumber"`
		// "1", "2", ... or "S1", "S2", ...
		AnidbEpisode string `json:"anidbEpisode"`
		// Episode number within the AniDB entry.
		// If special, number will be parsed to int.
		EpisodeNumber int `json:"episodeNumber"`
		// Absolute episode number across all seasons.
		AbsoluteEpisodeNumber int `json:"absoluteEpisodeNumber"`
		// Relative episode number within a TVDB season.
		TvdbRelativeEpisodeNumber int               `json:"tvdbRelativeEpisodeNumber"`
		Title                     map[string]string `json:"title"`
		AirDate                   string            `json:"airDate"`
		Overview                  string            `json:"overview"`
		Image                     string            `json:"image"`
	}
)

type SeriesMetadataRepository struct {
	logger *zerolog.Logger
}

func NewSeriesMetadata(anilistId int, animeListData *animelists.AnimeListFullData) (ret *SeriesMetadata, err error) {

	item, ok := animeListData.ItemsByAnilistID[anilistId]
	if !ok {
		return nil, fmt.Errorf("could not find anime with Anilist ID %d in offline database", anilistId)
	}

	if item.TheTvdbID == 0 {
		return nil, fmt.Errorf("could not find TVDB ID for anime with Anilist ID %d in offline database", anilistId)
	}

	ret = &SeriesMetadata{
		AnilistID:    anilistId,
		TvdbID:       item.TheTvdbID,
		AnidbID:      item.AnidbID,
		Titles:       make(map[string]string),           // to fill
		Episodes:     make(map[string]*EpisodeMetadata), // to fill
		EpisodeCount: 0,                                 // to fill
		SpecialCount: 0,                                 // to fill
		SeasonNumber: 0,                                 // to fill
	}

	return
}
