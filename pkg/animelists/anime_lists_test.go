package animelists

import (
	"github.com/5rahim/hibike/internal/testutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAnimeOfflineDatabaseResponse(t *testing.T) {

	bench := testutil.StartBenchmark("Fetching entire anime offline database")

	data, err := GetReducedAnimeOfflineDatabaseBytes()
	require.NoError(t, err)

	resp, err := GetReducedAnimeOfflineDatabase(data)
	require.NoError(t, err)

	bench.Print()

	t.Logf("Count: %d\n", resp.Count)

	var anilistId = 20912 // Hibike! Euphonium

	anime, found := resp.ItemsByAnilistID[anilistId]
	require.True(t, found)

	t.Logf("Anilist ID: %d\n", anime.AnilistID)
	t.Logf("\t\tAniDB ID: %d\n", anime.AnidbID)
}

func TestGetAnimeListsFull(t *testing.T) {

	bench := testutil.StartBenchmark("Fetching anime lists full")

	data, err := GetAnimeListFullBytes()
	require.NoError(t, err)

	resp, err := GetAnimeListFull(data)
	require.NoError(t, err)

	bench.Print()

	t.Logf("Count: %d\n", resp.Count)

	var anilistId = 20912 // Hibike! Euphonium

	anime, found := resp.ItemsByAnilistID[anilistId]
	require.True(t, found)

	t.Logf("Anilist ID: %d\n", anime.AnilistID)
	t.Logf("\t\tAniDB ID: %d\n", anime.AnidbID)
	t.Logf("\t\tTheTVDB ID: %d\n", anime.TheTvdbID)

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const animeOfflineDatabaseTestDataFile = "anime_offline_database.json"

func TestStoreAnimeOfflineDatabase(t *testing.T) {

	data, err := GetReducedAnimeOfflineDatabaseBytes()
	require.NoError(t, err)

	res, err := GetReducedAnimeOfflineDatabase(data)
	require.NoError(t, err)

	err = testutil.SaveTestDataJSON(animeOfflineDatabaseTestDataFile, res)
	require.NoError(t, err)

}

const animeListFullTestDataFile = "anime_list_full.json"

func TestStoreAnimeListsFull(t *testing.T) {

	data, err := GetAnimeListFullBytes()
	require.NoError(t, err)

	spew.Dump(string(data))

	res, err := GetAnimeListFull(data)
	require.NoError(t, err)

	err = testutil.SaveTestDataJSON(animeListFullTestDataFile, res)
	require.NoError(t, err)

}
