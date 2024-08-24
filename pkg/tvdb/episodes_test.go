package tvdb

import (
	testutil "github.com/5rahim/hibike/internal/testutil"
	"github.com/5rahim/hibike/pkg/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTVDB_FetchEpisodes(t *testing.T) {

	var seriesID = 289884 // Hibike! Euphonium

	tvdb := NewTVDB(&NewTVDBOptions{
		ApiKey: "",
		Logger: util.NewLogger(),
	})

	bench := testutil.StartBenchmark("Fetching episodes")

	show, err := tvdb.FetchShow(seriesID)
	require.NoError(t, err)

	absoluteSeason, found := show.GetAbsoluteSeason()
	require.True(t, found)

	t.Logf("Absolute Season ID: %+v\n", absoluteSeason.ID)

	seasonWithEpisodes, err := tvdb.FetchSeasonEpisodes(absoluteSeason)
	require.NoError(t, err)

	bench.Print()

	for _, episode := range seasonWithEpisodes.Episodes {
		t.Logf("Episode ID: %+v\n", episode.ExtendedSeasonEpisode.ID)
		t.Logf("\t\tSeason Number: %+v\n", episode.ExtendedSeasonEpisode.SeasonNumber)
		t.Logf("\t\tName: %+v\n", episode.GetEnglishTitle())
		t.Logf("\t\tOverview: %+v\n", episode.GetEnglishOverview())
		t.Logf("\t\tNumber: %+v\n", episode.ExtendedSeasonEpisode.Number)
		t.Logf("\t\tImage: %+v\n", episode.ExtendedSeasonEpisode.Image)
		t.Logf("\t\tLast Updated: %+v\n", episode.ExtendedSeasonEpisode.LastUpdated)
		t.Logf("\t\tName Translations: %+v\n", episode.ExtendedSeasonEpisode.NameTranslations)
		t.Logf("\t\tAired: %+v\n", episode.ExtendedSeasonEpisode.Aired)
		t.Logf("\t\tRuntime: %+v\n", episode.ExtendedSeasonEpisode.Runtime)
	}

}

func TestTVDB_FetchEpisodeTranslations(t *testing.T) {

	var episodeId int64 = 10382953 // Hibike! Euphonium S01E01

	tvdb := NewTVDB(&NewTVDBOptions{
		ApiKey: "",
		Logger: util.NewLogger(),
	})

	bench := testutil.StartBenchmark("Fetching episode translations")

	translations, err := tvdb.fetchEpisodeTranslations(episodeId, "eng")
	require.NoError(t, err)

	bench.Print()

	spew.Dump(translations)

}
