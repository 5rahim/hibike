package animelists

import (
	"github.com/5rahim/hibike/pkg/util/common"
	"strings"
)

// FindMappingAnimeByAnidbID returns the [MappingAnime] object for the given AniDB ID.
// This should be used to find the mapping for a specific Anilist anime entry since it is unique to each anime season.
//
//	<anime anidbid="10889" tvdbid="289884" defaulttvdbseason="1">
//	    <name>Hibike! Euphonium</name>
//	    <mapping-list>
//	        <mapping anidbseason="0" tvdbseason="0">;1-8;</mapping>
//	        <mapping anidbseason="0" tvdbseason="0" start="2" end="8" offset="-1"/>
//	    </mapping-list>
//	</anime>
func (m *MappingAnimeList) FindMappingAnimeByAnidbID(anidbId int) (ret *MappingAnime, found bool) {
	for _, item := range m.Anime {
		if common.ToIntMust(item.AnidbID) == anidbId {
			return item, true
		}
	}
	return nil, false
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GetTvdbID returns the corresponding TheTVDB ID as an integer.
// For movies, [MappingAnime.TvdbID] will be "movie", so this will return -1.
//
//	<anime anidbid="10889" tvdbid="289884" defaulttvdbseason="1">
func (a *MappingAnime) GetTvdbID() int {
	return common.ToIntOr(a.TvdbID, -1)
}

// GetDefaultTvdbSeason returns the corresponding TheTVDB season [MappingAnime.DefaultTvdbSeason].
// For one-off titles it will be "1" unless associated to a multi-episode series, in which case it will be "0".
// Series that span multiple seasons on TheTVDB.com may be marked as "a" if the absolute episode numbering is defined and matches AniDB.net.
//
//	<anime anidbid="10889" tvdbid="289884" defaulttvdbseason="1">
func (a *MappingAnime) GetDefaultTvdbSeason() string {
	return a.DefaultTvdbSeason
}

// IsTvdbSpecial returns true if [MappingAnime.DefaultTvdbSeason] is "0".
func (a *MappingAnime) IsTvdbSpecial() bool {
	return a.DefaultTvdbSeason == "0"
}

// IsTvdbAbsolute returns true if [MappingAnime.DefaultTvdbSeason] is "a".
func (a *MappingAnime) IsTvdbAbsolute() bool {
	return a.DefaultTvdbSeason == "a"
}

func (a *MappingAnime) IsTvdbEntry() bool {
	return common.ToIntOr(a.TvdbID, -1) != -1
}

// GetEpisodeOffset returns the number to add to each regular AniDB.net episode number
// to get the corresponding TheTVDB episode number in [MappingAnime.GetDefaultTvdbSeason].
// Not necessary if the episode numbers match up exactly. For special episodes and more complex situations the mapping-list element should be used.
func (a *MappingAnime) GetEpisodeOffset() (int, bool) {
	if a.EpisodeOffset == "" {
		return 0, false
	}
	return common.ToIntOr(a.EpisodeOffset, -1), true
}

//////////////////////////////////////////////

func (a *MappingAnime) GetAnidbID() int {
	return common.ToIntMust(a.AnidbID)
}

//////////////////////////////////////////////

// MappingData is parsed data.
//
//	e.g. The following tells us that this OVA has a regular entry on AniDB (i.e. season 1) but on TVDB it is in the main entry as a special (i.e. season 0).
//	<anime anidbid="905" tvdbid="81472" defaulttvdbseason="0" episodeoffset="11" imdbid="tt0142242">
//	    <name>Dragon Ball Z: Moetsukiro!! Nessen, Ressen, Chougekisen</name>
//	    <mapping-list>
//	        <mapping anidbseason="1" tvdbseason="0">;2-8;3-8;</mapping>
//	    </mapping-list>
//	 </anime>
//
//		Items: [ { AnidbSeason: 1, TvdbSeason: 0, AnidbToTvdb: { 2: [8], 3: [8] } } ]
//		EpisodeOffset: 11
type MappingData struct {
	Items []*MappingListDataItem
	// Number to add to each regular AniDB.net episode number to get the corresponding TheTVDB episode number in the [MappingAnime.DefaultTvdbSeason].
	// Not necessary if the episode numbers match up exactly.
	EpisodeOffset int
}

// MappingListDataItem represents a single mapping between AniDB and TheTVDB.
// [MappingListDataItem.AnidbToTvdb] is a map of [MappingListDataItem.AnidbSeason] episode numbers to [MappingListDataItem.TvdbSeason] episode numbers.
//
//	e.g. <mapping anidbseason="0" tvdbseason="0">;1-8;</mapping>
//	AnidbToTvdb: { AnidbSeason: 0, TvdbSeason: 0, AnidbToTvdb: { 1: [8] } }
//
//	e.g. <mapping anidbseason="0" tvdbseason="0" start="2" end="8" offset="-1"/>
//	AnidbToTvdb: { AnidbSeason: 0, TvdbSeason: 0, AnidbToTvdb: { 2: [1], 3: [2], 4: [3], 5: [4], 6: [5], 7: [6], 8: [7] } }
//
//	e.g. <mapping anidbseason="0" tvdbseason="1" start="1" end="7" offset="13"/>
//	AnidbToTvdb: { AnidbSeason: 0, TvdbSeason: 1, AnidbToTvdb: { 1: [14], 2: [15], 3: [16], 4: [17], 5: [18], 6: [19], 7: [20] } }
type MappingListDataItem struct {
	// 0 for specials, 1 for main series.
	// A main series can be mapped to multiple TheTVDB seasons if [MappingAnime.DefaultTvdbSeason] is "a".
	AnidbSeason int
	// 0 for specials, 1, 2, 3, ... for series' seasons.
	TvdbSeason  int
	AnidbToTvdb map[int][]int
}

// GetMappingData returns a list of [MappingListDataItem] parsed from [MappingAnime.MappingList].
func (a *MappingAnime) GetMappingData() (ret *MappingData, ok bool) {

	// If the anime is a movie, return an empty list
	if !a.IsTvdbEntry() {
		return nil, false
	}

	ret = &MappingData{
		Items:         make([]*MappingListDataItem, 0),
		EpisodeOffset: common.ToIntOr(a.EpisodeOffset, 0),
	}

	for _, mapping := range a.MappingList.Mapping {
		anidbSeason := common.ToIntMust(mapping.AnidbSeason)
		tvdbSeason := common.ToIntMust(mapping.TvdbSeason)
		anidbToTvdb := mapping.parseMapping()
		ret.Items = append(ret.Items, &MappingListDataItem{
			AnidbSeason: anidbSeason,
			TvdbSeason:  tvdbSeason,
			AnidbToTvdb: anidbToTvdb,
		})
	}
	return ret, true
}

// parseMapping
// Docs:
//
//	mapping-list - Used to map individual episodes between AniDB.net and TheTVDB.com (see below). Not necessary if episode numbers match up exactly within the season(s).
//	The mapping-list node consists of one or more mapping nodes with the following attributes:
//	anidbseason - The AniDB.net season (either 1 for regular episodes or 0 for specials).
//	tvdbseason - The corresponding TheTVDB.com season.
//	start - The first AniDB.net episode the offset applies to.
//	end - The last AniDB.net episode the offset applies to.
//	offset - The number to add to each AniDB.net episode between the start and end values.
//	The format for mapping individual episodes is: ;1-5;2-6;...; where the first number in each mapping is the AniDB.net episode number and the second is the corresponding TheTVDB.com episode number for the season specified. For a single AniDB.net episode to multiple TheTVDB.com episodes, you can link the latter with +. For example ;1-1+2; if the AniDB.net episode 1 contains both TheTVDB.com episodes 1 and 2. Conversely, multiple AniDB.net episodes simply map to the same TheTVDB.com episode. For example ;1-1;2-1; if AniDB.net episodes 1 and 2 are parts of TheTVDB.com episode 1. (;1+2-1; IS NOT VALID!)
//	AniDB trailer episodes with prefix T can be mapped by using episode numbers >= 201, T1 = 201. "Other" episodes can be mapped using episode numbers >= 401, O1 = 401.
//	The start, end, and offset attributes are not necessary if only individual episodes are being mapped. The offset and/or episodeoffset will be ignored in favour of an individual mapping.
//	Episodes on AniDB.net that don't match anything on TheTVDB.com are mapped to 0 if and only if there's a conflict.
func (m *Mapping) parseMapping() (anidbToTvdb map[int][]int) {
	anidbToTvdb = make(map[int][]int)

	rangeMap, ok := m.parseRange()
	if ok {
		for anidb, tvdb := range rangeMap {
			anidbToTvdb[anidb] = tvdb
		}
		return
	}

	valueMap, ok := m.parseValue()
	if ok {
		for anidb, tvdb := range valueMap {
			anidbToTvdb[anidb] = tvdb
		}
		return
	}

	return
}

// parseRange
//
//	e.g. <mapping anidbseason="1" tvdbseason="1" start="1" end="20"/>
//	output: map[1:[1] 2:[2] 3:[3] 4:[4] 5:[5] ... 19:[19] 20:[20]]
func (m *Mapping) parseRange() (anidbToTvdb map[int][]int, ok bool) {

	if m.Start == "" || m.End == "" {
		return nil, false
	}

	anidbToTvdb = make(map[int][]int)

	// Parse the range
	start := common.ToIntOr(m.Start, -1)
	end := common.ToIntOr(m.End, -1)
	offset := common.ToIntOr(m.Offset, 0)

	if start == -1 || end == -1 {
		return nil, false
	}

	// Add the range to the list
	for i := start; i <= end; i++ {
		appendToIntSliceMap(anidbToTvdb, i, i+offset)
	}

	return anidbToTvdb, true
}

// parseValue
//
//	e.g. <mapping anidbseason="0" tvdbseason="0">;1-2;2-3;3-4+5;4-0;5-0;</mapping>
//	output: map[1:[2] 2:[3] 3:[4 5] 4:[0] 5:[0]]
func (m *Mapping) parseValue() (anidbToTvdb map[int][]int, ok bool) {
	if m.Value == "" {
		return nil, false
	}

	anidbToTvdb = make(map[int][]int)

	// Parse the value
	value := m.Value
	pairs := strings.Split(value, ";")

	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		parts := strings.Split(pair, "-")
		anidb := common.ToIntMust(parts[0])
		tvdbStr := parts[1]
		// Check if the tvdb episode is composed of multiple episodes
		if strings.Contains(tvdbStr, "+") {
			tvdbParts := strings.Split(tvdbStr, "+")
			for _, tvdb := range tvdbParts {
				appendToIntSliceMap(anidbToTvdb, anidb, common.ToIntMust(tvdb))
			}
		} else {
			appendToIntSliceMap(anidbToTvdb, anidb, common.ToIntMust(tvdbStr))
		}
	}

	return anidbToTvdb, true
}

func appendToIntSliceMap(m map[int][]int, key, value int) {
	if _, ok := m[key]; ok {
		m[key] = append(m[key], value)
	} else {
		m[key] = []int{value}
	}
}
