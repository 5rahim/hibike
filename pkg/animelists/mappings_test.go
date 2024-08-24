package animelists

import (
	"github.com/5rahim/hibike/internal/testutil"
	"github.com/5rahim/hibike/pkg/util/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMappingAnime_GetMappingData(t *testing.T) {

	var mappings *MappingAnimeList
	err := testutil.LoadTestDataXML(mappingMasterTestDataFile, &mappings)
	require.NoError(t, err)

	tests := []struct {
		name            string
		anidbId         int
		expectedTvdbID  int
		matchesAbsolute bool
	}{
		{
			name:            "Hibike! Euphonium 2",
			anidbId:         11746,
			expectedTvdbID:  289884,
			matchesAbsolute: false,
		},
		{
			name:            "One Piece",
			anidbId:         69,
			expectedTvdbID:  81797,
			matchesAbsolute: true, // One Piece has a single entry on AniDB (absolute numbering) & TVDB also has absolute numbering
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mapping, found := mappings.FindMappingAnimeByAnidbID(tt.anidbId)
			require.True(t, found)

			require.Equal(t, tt.expectedTvdbID, mapping.GetTvdbID())

			mappingData, ok := mapping.GetMappingData()
			require.Truef(t, ok, "failed to get mapping data for %s", tt.name)
			require.NotNil(t, mappingData, "mapping data is nil")

			for _, item := range mappingData.Items {
				t.Logf("AniDB Season %d, TVDB Season: %d\n", item.AnidbSeason, item.TvdbSeason)
				for anidbEp, tvdbEp := range item.AnidbToTvdb {
					t.Logf("\t\tAniDB Ep %d -> TVDB Ep %+v\n", anidbEp, tvdbEp)
				}
			}

			require.Equal(t, tt.matchesAbsolute, mapping.IsTvdbAbsolute())
		})
	}

}

func TestMappingOperations(t *testing.T) {

	var mappings *MappingAnimeList
	err := testutil.LoadTestDataXML(mappingMasterTestDataFile, &mappings)
	require.NoError(t, err)

	tests := []struct {
		name    string
		anidbId int
	}{
		{
			name:    "Hibike! Euphonium",
			anidbId: 10889,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			item, found := mappings.FindMappingAnimeByAnidbID(tt.anidbId)
			require.True(t, found)

			t.Logf("AniDB ID: %s\n", item.AnidbID)
			t.Logf("\t\tTVDB ID: %s\n", item.TvdbID)
			t.Logf("\t\tTVDB Default Season: %s\n", item.DefaultTvdbSeason)
			t.Logf("\t\tMapping: %+v\n", item.MappingList.Mapping)
		})
	}

}

func TestGetMappingMaster(t *testing.T) {

	data, err := GetMappingMasterBytes()
	require.NoError(t, err)

	res, err := GetMappingMaster(data)
	require.NoError(t, err)

	t.Logf("Count: %d\n", len(res.Anime))

	for _, item := range res.Anime {
		if common.ToIntMust(item.AnidbID) != 10889 {
			continue
		}

		t.Logf("AniDB ID: %s\n", item.AnidbID)
		t.Logf("\t\tTVDB ID: %s\n", item.TvdbID)
		t.Logf("\t\tTVDB Default Season: %s\n", item.DefaultTvdbSeason)
		t.Logf("\t\tMapping: %+v\n", item.MappingList.Mapping)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const mappingMasterTestDataFile = "animelists_mapping_master.xml"

func TestStoreMappingMaster(t *testing.T) {

	data, err := GetMappingMasterBytes()
	require.NoError(t, err)

	res, err := GetMappingMaster(data)
	require.NoError(t, err)

	err = testutil.SaveTestDataXML(mappingMasterTestDataFile, res)
	require.NoError(t, err)

}
