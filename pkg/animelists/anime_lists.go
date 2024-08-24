package animelists

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const (
	AnimeListsOfflineDatabaseReducedURL = "https://raw.githubusercontent.com/Fribb/anime-lists/master/anime-offline-database-reduced.json"
	AnimeListsFullURL                   = "https://raw.githubusercontent.com/Fribb/anime-lists/master/anime-list-full.json"
)

type (
	ReducedAnimeOfflineDatabaseData struct {
		Items            []*ReducedAnimeOfflineDatabaseItem
		ItemsByAnilistID map[int]*ReducedAnimeOfflineDatabaseItem
		Count            int
	}

	ReducedAnimeOfflineDatabaseItem struct {
		AnidbID   int `json:"anidb_id"`
		AnilistID int `json:"anilist_id"`
		KitsuID   int `json:"kitsu_id"`
		MalID     int `json:"mal_id"`
	}

	/////////////////////////////////////////////////////////////

	AnimeListFullData struct {
		Items            []*AnimeListFullItem
		ItemsByAnilistID map[int]*AnimeListFullItem
		Count            int
	}

	AnimeListFullItem struct {
		AnidbID      int         `json:"anidb_id"`
		AnilistID    int         `json:"anilist_id"`
		KitsuID      int         `json:"kitsu_id"`
		TheTvdbID    int         `json:"thetvdb_id"`
		TheMovieDbID interface{} `json:"themoviedb_id"` // Can be int or string
		MalID        int         `json:"mal_id"`
		// We don't need the rest
	}
)

func GetReducedAnimeOfflineDatabaseBytes() (resp []byte, err error) {
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

func GetReducedAnimeOfflineDatabase(data []byte) (resp *ReducedAnimeOfflineDatabaseData, err error) {
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

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetAnimeListFullBytes() (resp []byte, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", AnimeListsFullURL, nil)
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

func GetAnimeListFull(data []byte) (resp *AnimeListFullData, err error) {
	var items []*AnimeListFullItem
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("animelists: failed to decode data: %w", err)
	}

	itemsByAnilistID := make(map[int]*AnimeListFullItem)
	for _, item := range items {
		if item.AnilistID == 0 {
			continue
		}
		itemsByAnilistID[item.AnilistID] = item
	}

	return &AnimeListFullData{
		Items:            items,
		ItemsByAnilistID: itemsByAnilistID,
		Count:            len(items),
	}, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (i *ReducedAnimeOfflineDatabaseData) GetItemByAnilistID(anilistID int) (ret *ReducedAnimeOfflineDatabaseItem, found bool) {
	ret, found = i.ItemsByAnilistID[anilistID]
	return
}
