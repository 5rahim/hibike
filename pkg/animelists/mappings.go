package animelists

import (
	"encoding/xml"
	"io"
	"net/http"
)

const (
	MappingMasterURL = "https://raw.githubusercontent.com/Anime-Lists/anime-lists/master/anime-list-master.xml"
)

type (
	MappingAnimeList struct {
		XMLName *xml.Name       `xml:"anime-list"`
		Anime   []*MappingAnime `xml:"anime"`
	}

	MappingAnime struct {
		XMLName           *xml.Name      `xml:"anime"`
		AnidbID           string         `xml:"anidbid,attr"`
		TvdbID            string         `xml:"tvdbid,attr"`
		DefaultTvdbSeason string         `xml:"defaulttvdbseason,attr"`
		EpisodeOffset     string         `xml:"episodeoffset,attr,omitempty"`
		TmdbID            string         `xml:"tmdbid,attr,omitempty"`
		Name              string         `xml:"name"`
		MappingList       *MappingList   `xml:"mapping-list"`
		Before            *MappingBefore `xml:"before"`
	}

	MappingList struct {
		XMLName *xml.Name  `xml:"mapping-list"`
		Mapping []*Mapping `xml:"mapping"`
	}

	MappingBefore struct {
		XMLName *xml.Name `xml:"before"`
		Value   string    `xml:",chardata"`
	}

	Mapping struct {
		XMLName     xml.Name `xml:"mapping"`
		AnidbSeason string   `xml:"anidbseason,attr"`
		TvdbSeason  string   `xml:"tvdbseason,attr"`
		Value       string   `xml:",chardata"`
		Start       string   `xml:"start,attr,omitempty"`
		End         string   `xml:"end,attr,omitempty"`
		Offset      string   `xml:"offset,attr,omitempty"`
	}
)

func GetMappingMasterBytes() (ret []byte, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", MappingMasterURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret, err = io.ReadAll(res.Body)

	return ret, err
}

func GetMappingMaster(data []byte) (resp *MappingAnimeList, err error) {
	var ret MappingAnimeList
	if err = xml.Unmarshal(data, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
