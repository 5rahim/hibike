package tvdb

import (
	"fmt"
	"github.com/goccy/go-json"
	"sync"
)

type (
	SeasonWithEpisodes struct {
		ExtendedSeason *ExtendedSeason `json:"season"`
		Episodes       []*Episode      `json:"episodes"`
	}

	Episode struct {
		ExtendedSeasonEpisode *ExtendedSeasonEpisode `json:"episode"`
		Translation           *Translation           `json:"translation"`
	}
)

func (tvdb *TVDB) FetchSeasonEpisodes(season *ExtendedSeason) (ret *SeasonWithEpisodes, err error) {

	tvdb.logger.Debug().Msg("tvdb: Fetching all possible episodes")

	ret = &SeasonWithEpisodes{
		ExtendedSeason: season,
		Episodes:       make([]*Episode, 0),
	}

	tvdb.logger.Debug().Int64("seasonId", season.ID).Msg("tvdb: Fetching episodes for season")

	// Fetch season metadata
	resp, err := tvdb.doRequest(fmt.Sprintf("%s/seasons/%d/extended", ApiUrl, season.ID), nil)
	if err != nil {
		return ret, fmt.Errorf("tvdb: could not fetch season metadata: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var data ExtendedSeasonResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		tvdb.logger.Error().Int64("seasonId", season.ID).Err(err).Msg("tvdb: Could not decode response")
		return
	}

	if data.Data == nil || data.Data.Episodes == nil {
		tvdb.logger.Warn().Int64("seasonId", season.ID).Msg("tvdb: Could not find episodes for season")
		return
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(data.Data.Episodes))
	for _, episode := range data.Data.Episodes {
		go func(episode *ExtendedSeasonEpisode) {
			defer wg.Done()

			// Fetch episode translation
			translation, _ := tvdb.fetchEpisodeTranslations(episode.ID, "eng")

			mu.Lock()
			ret.Episodes = append(ret.Episodes, &Episode{
				ExtendedSeasonEpisode: episode,
				Translation:           translation,
			})
			mu.Unlock()
		}(episode)
	}
	wg.Wait()

	return
}

func (tvdb *TVDB) fetchEpisodeTranslations(episodeID int64, lang string) (ret *Translation, err error) {
	resp, err := tvdb.doRequest(fmt.Sprintf("%s/episodes/%d/translations/%s", ApiUrl, episodeID, lang), nil)
	if err != nil {
		return nil, fmt.Errorf("tvdb: could not fetch episode translation: %w", err)
	}
	defer resp.Body.Close()

	var res *TranslationResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("tvdb: could not decode episode translation response: %w", err)
	}

	if res.Status != "success" {
		return nil, fmt.Errorf("tvdb: episode translation response status is not success")
	}

	return res.Data, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (e *Episode) GetEnglishTitle() string {
	if e.Translation == nil {
		return ""
	}
	return e.Translation.Name
}

func (e *Episode) GetEnglishOverview() string {
	if e.Translation == nil {
		return ""
	}

	return e.Translation.Overview
}
