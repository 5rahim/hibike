package tvdb

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strings"
)

type (
	TVDB struct {
		apiKey       string
		client       *http.Client
		currentToken string // Hydrated by getTokenWithKey
		logger       *zerolog.Logger
	}

	NewTVDBOptions struct {
		ApiKey string
		Logger *zerolog.Logger
	}

	FilterEpisodeMediaInfo struct {
		Year           *int
		Month          *int
		TotalEp        int // from anizip
		AbsoluteOffset int // from anizip
	}
)

func NewTVDB(opts *NewTVDBOptions) (ret *TVDB) {
	ret = &TVDB{
		apiKey: opts.ApiKey,
		client: &http.Client{},
		logger: opts.Logger,
	}

	// Get token
	_, err := ret.getTokenWithTries()
	if err != nil {
		return
	}

	return
}

//
//func (tvdb *TVDB) FetchSeriesEpisodes(id int, filter FilterEpisodeMediaInfo) (res []*Episode, err error) {
//	// Get token
//	_, err = tvdb.getTokenWithTries()
//	if err != nil {
//		return nil, err
//	}
//
//	// Fetch seasons
//	seasons, err := tvdb.FetchShow(id)
//	if err != nil {
//		return nil, err
//	}
//
//	// Fetch episodes
//	_, err = tvdb.FetchSeasonEpisodes(seasons, true)
//	if err != nil {
//		return nil, err
//	}
//
//	// Filter episodes
//	//episodesF, err = tvdb.filterEpisodes(episodesF, filter, true)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//// Convert episodes
//	//res = make([]*Episode, len(episodesF), len(episodesF))
//	//for i, e := range episodesF {
//	//	res[i] = &Episode{
//	//		ID:      e.ID,
//	//		Image:   e.Image,
//	//		Number:  int(e.Number),
//	//		AiredAt: e.Aired,
//	//	}
//	//}
//
//	tvdb.logger.Debug().Int("id", id).Int("episodes", len(res)).Msg("tvdb: Found episodes")
//
//	return
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tvdb *TVDB) doRequest(url string, body io.Reader) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tvdb.currentToken))

	return tvdb.client.Do(req)
}

// getTokenWithTries tries to get a token with all available API keys.
// If an API key is provided in the options, it will be tried first.
func (tvdb *TVDB) getTokenWithTries() (token string, err error) {

	if tvdb.apiKey != "" {
		token, err := tvdb.getTokenWithKey(tvdb.apiKey)
		if err == nil {
			return token, nil
		}
	}

	for _, key := range ApiKeys {
		token, err = tvdb.getTokenWithKey(key)
		if err != nil {
			continue
		}
		return token, nil
	}

	return "", fmt.Errorf("could not get authentication token")
}

// getTokenWithKey gets a token with a specific API key.
func (tvdb *TVDB) getTokenWithKey(key string) (token string, err error) {
	req, err := tvdb.client.Post(fmt.Sprintf("%s/login", ApiUrl), "application/json", strings.NewReader(fmt.Sprintf(`{"apikey":"%s"}`, key)))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	b, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(b, &res); err != nil {
		return "", err
	}

	data, ok := res["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("could not get token")
	}

	token, ok = data["token"].(string)
	if !ok {
		return "", fmt.Errorf("could not get token")
	}

	tvdb.currentToken = token

	return token, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (tvdb *TVDB) filterEpisodes(
//	episodes []*ExtendedSeasonEpisode,
//	mediaInfo FilterEpisodeMediaInfo,
//	isAbsolute bool,
//) (res []*ExtendedSeasonEpisode, err error) {
//
//	defer func() {
//		if r := recover(); r != nil {
//			tvdb.logger.Error().Interface("r", r).Msg("tvdb: Panic in filterEpisodes")
//		}
//	}()
//
//	res = episodes
//
//	if isAbsolute {
//
//		minEpNum := int64(mediaInfo.AbsoluteOffset + 1)
//		maxEpNum := int64(mediaInfo.AbsoluteOffset + mediaInfo.TotalEp)
//
//		filteredEpisodes := make([]*ExtendedSeasonEpisode, 0)
//		for _, episode := range episodes {
//			if episode.Number >= minEpNum && episode.Number <= maxEpNum {
//				filteredEpisodes = append(filteredEpisodes, episode)
//			}
//		}
//
//		if len(filteredEpisodes) == 0 {
//			return nil, errors.New("no episodes found")
//		}
//
//		if mediaInfo.AbsoluteOffset > 0 {
//			allAbsolute := true
//			for _, e := range filteredEpisodes {
//				if e.Number < int64(mediaInfo.TotalEp) {
//					allAbsolute = false
//					break
//				}
//			}
//
//			// Normalize the episodes
//			if allAbsolute {
//				for i, e := range filteredEpisodes {
//					filteredEpisodes[i].Number = e.Number - int64(mediaInfo.AbsoluteOffset)
//				}
//			}
//		}
//
//		res = filteredEpisodes
//		return res, nil
//	}
//
//	type kEpisode struct {
//		episode *ExtendedSeasonEpisode
//		factor  float64
//	}
//
//	// If we find episodes that are over the total episode count, we need to apply the offset
//	if mediaInfo.TotalEp > 0 {
//
//		// If the media has a month, we don't need to apply the offset
//		// just filter the episodes based on the month
//		if mediaInfo.Month != nil && *mediaInfo.Month > 0 {
//
//			mediaFactor := float64(*mediaInfo.Year) + (float64(*mediaInfo.Month) / 100) // e.g. 2021.05, 2021.12
//
//			// Filter episodes
//			kEpisodes := make([]*kEpisode, 0)
//
//			for _, episode := range episodes {
//				if episode.Aired == "" || episode.Year == "" {
//					continue
//				}
//				airedParts := strings.Split(episode.Aired, "-")
//				episodeMonth, _ := strconv.Atoi(airedParts[1])
//				episodeDay, _ := strconv.Atoi(airedParts[2])
//				episodeYear, _ := strconv.Atoi(episode.Year)
//				episodeFactor := float64(episodeYear) + (float64(episodeMonth) / 100) + (float64(episodeDay) / 10000)
//				// If the episode aired AFTER the month/year, we can include it
//				//spew.Printf("(%d) %s %s %d %d %f %f\n", episode.Number, episode.Aired, episode.Year, *mediaInfo.Year, *mediaInfo.Month, episodeFactor, mediaFactor)
//				if episodeYear > *mediaInfo.Year || (episodeYear == *mediaInfo.Year && episodeMonth >= *mediaInfo.Month) {
//					kEpisodes = append(kEpisodes, &kEpisode{
//						episode: episode,
//						factor:  episodeFactor,
//					})
//				}
//			}
//
//			// Sort episodes by factor (ascending)
//			for i := 0; i < len(kEpisodes); i++ {
//				for j := i + 1; j < len(kEpisodes); j++ {
//					if kEpisodes[i].factor > kEpisodes[j].factor {
//						kEpisodes[i], kEpisodes[j] = kEpisodes[j], kEpisodes[i]
//					}
//				}
//			}
//
//			//spew.Dump(mediaInfo)
//			//for _, kEpisode := range kEpisodes {
//			//	spew.Printf("(%d) %s %f %f\n", kEpisode.episode.Number, kEpisode.episode.Aired, kEpisode.factor, mediaFactor)
//			//}
//
//			// Keep episodes that are after the media factor but whose number is less than the total episode count
//			filteredEpisodes := make([]*ExtendedSeasonEpisode, 0)
//			addedAiredDates := make(map[string]*kEpisode) // Aired date -> Episode
//			count := 0
//			for _, kEpisode := range kEpisodes {
//				if kEpisode.factor >= mediaFactor {
//					if count < mediaInfo.TotalEp {
//
//						prev, ok := addedAiredDates[kEpisode.episode.Aired]
//
//						// episodesAiredSameDay
//						if ok && prev.episode.Number != kEpisode.episode.Number {
//							diff := math.Abs(float64(prev.episode.Number) - float64(kEpisode.episode.Number))
//							if diff < 12 {
//								filteredEpisodes = append(filteredEpisodes, kEpisode.episode)
//								addedAiredDates[kEpisode.episode.Aired] = kEpisode
//								continue
//							}
//						}
//						if !ok || (ok && prev.episode.Image == "") {
//							filteredEpisodes = append(filteredEpisodes, kEpisode.episode)
//							addedAiredDates[kEpisode.episode.Aired] = kEpisode
//							count++
//						}
//					}
//				}
//			}
//
//			if mediaInfo.AbsoluteOffset > 0 {
//				allAbsolute := true
//				for _, e := range filteredEpisodes {
//					if e.Number < int64(mediaInfo.TotalEp) {
//						allAbsolute = false
//						break
//					}
//				}
//
//				// Normalize the episodes
//				if allAbsolute {
//					for i, e := range filteredEpisodes {
//						filteredEpisodes[i].Number = e.Number - int64(mediaInfo.AbsoluteOffset)
//					}
//				}
//			}
//
//			//println("----------------")
//			//
//			//for _, kEpisode := range filteredEpisodes {
//			//	spew.Printf("(%d) %s\n", kEpisode.Number, kEpisode.Aired)
//			//}
//
//			if len(filteredEpisodes) == 0 {
//				return nil, errors.New("no episodes found")
//			}
//
//			res = filteredEpisodes
//		}
//
//	} else {
//
//		return nil, errors.New("could not filter episodes")
//
//	}
//
//	return res, nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
