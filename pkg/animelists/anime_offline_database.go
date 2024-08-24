package animelists

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const (
	AnimeOfflineDatabaseURL = "https://raw.githubusercontent.com/manami-project/anime-offline-database/master/anime-offline-database-minified.json"
)

type (
	AnimeOfflineDatabaseData struct {
		Items            []*AnimeData
		ItemsByAnilistID map[int]*ReducedAnimeOfflineDatabaseItem
		Count            int
	}

	AnimeDatabase struct {
		License    License      `json:"license"`
		Repository string       `json:"repository"`
		LastUpdate string       `json:"lastUpdate"`
		Data       []*AnimeData `json:"data"`
	}

	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	AnimeData struct {
		Sources      []string    `json:"sources"`
		Title        string      `json:"title"`
		Type         string      `json:"type"`
		Episodes     int         `json:"episodes"`
		Status       string      `json:"status"`
		AnimeSeason  AnimeSeason `json:"animeSeason"`
		Picture      string      `json:"picture"`
		Thumbnail    string      `json:"thumbnail"`
		Synonyms     []string    `json:"synonyms"`
		RelatedAnime []string    `json:"relatedAnime"`
		Tags         []string    `json:"tags"`
	}

	AnimeSeason struct {
		Season string `json:"season"`
		Year   int    `json:"year,omitempty"`
	}
)

func GetAnimeOfflineDatabaseBytes() (resp []byte, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", AnimeListsOfflineDatabaseReducedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("animelists: failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("animelists: failed to make request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("animelists: failed to read response: %w", err)
	}

	return data, nil
}

func GetAnimeOfflineDatabase(data []byte) (resp *ReducedAnimeOfflineDatabaseData, err error) {
	var items []*ReducedAnimeOfflineDatabaseItem
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("animelists: failed to decode data: %w", err)
	}

	itemsByAnilistID := make(map[int]*ReducedAnimeOfflineDatabaseItem)
	for _, item := range items {
		if item.AnilistID == 0 {
			continue
		}
		itemsByAnilistID[item.AnilistID] = item
	}

	return &ReducedAnimeOfflineDatabaseData{
		Items:            items,
		ItemsByAnilistID: itemsByAnilistID,
		Count:            len(items),
	}, nil
}
