package tvdb

import (
	"cmp"
	"github.com/5rahim/hibike/internal/testutil"
	"github.com/5rahim/hibike/pkg/util"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func TestTVDB_Show(t *testing.T) {

	var seriesID = 289884 // Hibike! Euphonium 1-3

	tvdb := NewTVDB(&NewTVDBOptions{
		ApiKey: "",
		Logger: util.NewLogger(),
	})

	bench := testutil.StartBenchmark("Fetching entire show")

	// Fetch the show
	show, err := tvdb.FetchShow(seriesID)
	require.NoError(t, err)

	// Get the absolute season
	absoluteSeason, found := show.GetAbsoluteSeason()
	require.True(t, found)

	t.Logf("Absolute Season ID: %+v\n", absoluteSeason.ID)

	// Fetch the season episodes
	seasonWithEpisodes, err := tvdb.FetchSeasonEpisodes(absoluteSeason)
	require.NoError(t, err)

	slices.SortStableFunc(seasonWithEpisodes.Episodes, func(i, j *Episode) int {
		return cmp.Compare(i.ExtendedSeasonEpisode.Number, j.ExtendedSeasonEpisode.Number)
	})

	bench.Print()

	for _, episode := range seasonWithEpisodes.Episodes {
		t.Logf("Episode ID: %+v\n", episode.ExtendedSeasonEpisode.ID)
		t.Logf("\t\tSeason Number: %+v\n", episode.ExtendedSeasonEpisode.SeasonNumber)
		t.Logf("\t\tName: %+v\n", episode.GetEnglishTitle())
		t.Logf("\t\tOverview: %+v\n", episode.GetEnglishOverview())
		t.Logf("\t\tNumber: %+v\n", episode.ExtendedSeasonEpisode.Number)
		t.Logf("\t\tImage: %+v\n", episode.ExtendedSeasonEpisode.Image)
		t.Logf("\t\tDescription: %+v\n", episode.Translation.Overview)
		t.Logf("\t\tLast Updated: %+v\n", episode.ExtendedSeasonEpisode.LastUpdated)
		t.Logf("\t\tName Translations: %+v\n", episode.ExtendedSeasonEpisode.NameTranslations)
	}

}

func TestTVDB_FetchSeasons(t *testing.T) {

	var seriesID = 289884 // Hibike! Euphonium 1-3

	tvdb := NewTVDB(&NewTVDBOptions{
		ApiKey: "",
		Logger: util.NewLogger(),
	})

	show, err := tvdb.FetchShow(seriesID)
	require.NoError(t, err)

	for _, season := range show.Seasons {
		t.Logf("Season ID: %+v\n", season.ID)
		t.Logf("\t\tNumber: %+v\n", season.Number)
		t.Logf("\t\tName: %+v\n", season.Type.Name)
		t.Logf("\t\tType: %+v\n", season.Type.Type)
		t.Logf("\t\tType: %+v\n", season.Type.Type)
		t.Logf("\t\tLast Updated: %+v\n", season.LastUpdated)
	}

}
