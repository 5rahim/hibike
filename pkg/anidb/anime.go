package anidb

import (
	"fmt"
	"github.com/5rahim/hibike/pkg/util/bypass"
	"github.com/5rahim/hibike/pkg/util/common"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gocolly/colly/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	AnimeURL = "https://anidb.net/anime/"
)

const (
	EpisodeTypeMain    EpisodeType = "main"
	EpisodeTypeSpecial EpisodeType = "special"
	EpisodeTypeTheme   EpisodeType = "theme"
	EpisodeTypeOther   EpisodeType = "other"
)

type (
	EpisodeType string

	AnimeData struct {
		ID                  int
		MainEpisodes        map[string]*Episode `json:"mainEpisodes"`
		SpecialEpisodes     map[string]*Episode `json:"specialEpisodes"`
		Themes              map[string]*Episode `json:"themes"`
		OtherEpisodes       map[string]*Episode `json:"otherEpisodes"`
		MainEpisodeCount    int                 `json:"mainEpisodeCount"`
		SpecialEpisodeCount int                 `json:"specialEpisodeCount"`
		ThemeCount          int                 `json:"themeCount"`
		OPEDCount           int                 `json:"opedEpisodeCount"`
	}

	Episode struct {
		ID int `json:"id"`
		// "1", "2", "3", "S1", "S2", "S3", ...
		Episode string `json:"number"`
		Number  int    `json:"numberInt"`
		Title   string `json:"title"`
		Runtime int    `json:"runtime"`
		// "YYYY-MM-DD"
		AirDate string      `json:"airDate"`
		Type    EpisodeType `json:"type"`
	}
)

func ScrapeAnime(id int) (ret *AnimeData, err error) {

	cl := &http.Client{
		Timeout: 60 * time.Second,
	}
	cl.Transport = bypass.AddCloudFlareByPass(cl.Transport)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", browser.MacOSX())
	})

	var episodes []*Episode

	c.OnHTML("tr[data-anidb-eid]", func(e *colly.HTMLElement) {
		eid := e.Attr("data-anidb-eid")
		if eid == "" {
			return
		}

		_eid, _ := strconv.Atoi(eid)

		ep := &Episode{
			ID: _eid,
		}

		e.ForEach("td", func(i int, e *colly.HTMLElement) {
			switch i {
			case 0:
				ep.Episode = strings.TrimSpace(e.ChildText("abbr"))
			case 1:
				ep.Title = strings.TrimSpace(e.ChildText("label"))
			case 2:
				ep.Runtime, err = strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(e.Text), "m"))
				if err != nil {
					ep.Runtime = 0
				}
			case 3:
				ep.AirDate = strings.TrimSpace(e.Attr("content"))
			}
		})

		episodes = append(episodes, ep)
	})

	tries := 0

	c.OnError(func(r *colly.Response, err error) {
		tries++
		if tries < 2 {
			err = r.Request.Retry()
			return
		} else {
			err = fmt.Errorf("anidb: failed to scrape anime %d: %w", id, err)
		}
	})

	err = c.Visit(fmt.Sprintf("%s%d", AnimeURL, id))
	if err != nil {
		return nil, err
	}

	ret = &AnimeData{
		ID:                  id,
		MainEpisodes:        make(map[string]*Episode),
		SpecialEpisodes:     make(map[string]*Episode),
		OtherEpisodes:       make(map[string]*Episode),
		Themes:              make(map[string]*Episode),
		MainEpisodeCount:    0,
		SpecialEpisodeCount: 0,
	}

	for _, ep := range episodes {
		if unicode.IsDigit(rune(ep.Episode[0])) {
			ret.MainEpisodeCount++
			ep.Number = common.ToIntOr(ep.Episode, -1)
			ep.Type = EpisodeTypeMain
			ret.MainEpisodes[ep.Episode] = ep
		} else if strings.HasPrefix(ep.Episode, "S") {
			ret.SpecialEpisodeCount++
			ep.Number = common.ToIntOr(strings.TrimPrefix(ep.Episode, "S"), -1)
			ep.Type = EpisodeTypeSpecial
			ret.SpecialEpisodes[ep.Episode] = ep
		} else if strings.HasPrefix(ep.Episode, "OP") || strings.HasPrefix(ep.Episode, "ED") {
			ret.ThemeCount++
			ep.Number = common.ExtractFirstIntOr(ep.Episode, -1)
			ep.Type = EpisodeTypeTheme
			ret.Themes[ep.Episode] = ep
		} else {
			ep.Number = -1
			ep.Type = EpisodeTypeOther
			ret.OtherEpisodes[ep.Episode] = ep
		}
	}

	return ret, nil
}
