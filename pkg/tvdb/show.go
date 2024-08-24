package tvdb

import (
	"fmt"
	"github.com/goccy/go-json"
	"time"
)

type (
	Show struct {
		ID      int64             `json:"id"`
		Seasons []*ExtendedSeason `json:"seasons"`
	}
)

func (tvdb *TVDB) FetchShow(id int) (res *Show, err error) {

	start := time.Now()
	tvdb.logger.Debug().Int("id", id).Msg("tvdb: Fetching seasons")

	// Fetch series metadata
	resp, err := tvdb.doRequest(fmt.Sprintf("%s/series/%d/extended", ApiUrl, id), nil)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	// Parse response
	var data ExtendedSeriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		tvdb.logger.Error().Err(err).Msg("tvdb: Could not decode response")
		return res, err
	}

	if data.Data == nil || data.Data.Seasons == nil {
		tvdb.logger.Error().Msg("tvdb: Could not find seasons")
		return res, fmt.Errorf("could not find seasons")
	}

	tvdb.logger.Debug().Int("id", id).Msgf("tvdb: Fetched seasons in %dms", time.Since(start).Milliseconds())

	res = &Show{
		ID:      int64(id),
		Seasons: data.Data.Seasons,
	}

	return res, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (show *Show) GetAbsoluteSeason() (*ExtendedSeason, bool) {
	for _, s := range show.Seasons {
		if s.IsAbsolute() {
			return s, true
		}
	}
	return nil, false
}

func (show *Show) GetSpecialsAndMoviesSeason() (*ExtendedSeason, bool) {
	for _, s := range show.Seasons {
		if s.IsSpecialsAndMovies() {
			return s, true
		}
	}
	return nil, false
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
