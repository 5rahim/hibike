package anidb

import (
	"github.com/5rahim/hibike/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScrapeAnime(t *testing.T) {

	bench := testutil.StartBenchmark("Scraping anime")

	data, err := ScrapeAnime(11746) // Hibike! Euphonium 2
	require.NoError(t, err)

	bench.Print()

	t.Logf("Main episodes: %d\n", len(data.MainEpisodes))

	for _, ep := range data.MainEpisodes {
		t.Logf("Episode ID: %d\n", ep.ID)
		t.Logf("\t\tEpisode: %s\n", ep.Episode)
		t.Logf("\t\tNumber: %d\n", ep.Number)
		t.Logf("\t\tTitle: %s\n", ep.Title)
		t.Logf("\t\tAirDate: %s\n", ep.AirDate)
		t.Logf("\t\tType: %s\n", ep.Type)
	}

	require.Equal(t, 13, len(data.MainEpisodes))

	t.Logf("\n\nSpecial episodes: %d\n", len(data.SpecialEpisodes))

	for _, ep := range data.SpecialEpisodes {
		t.Logf("Episode ID: %d\n", ep.ID)
		t.Logf("\t\tEpisode: %s\n", ep.Episode)
		t.Logf("\t\tNumber: %d\n", ep.Number)
		t.Logf("\t\tTitle: %s\n", ep.Title)
		t.Logf("\t\tAirDate: %s\n", ep.AirDate)
		t.Logf("\t\tType: %s\n", ep.Type)
	}

	require.Equal(t, 9, len(data.SpecialEpisodes))

	t.Logf("\n\nThemes: %d\n", len(data.Themes))

	for _, ep := range data.Themes {
		t.Logf("Episode ID: %d\n", ep.ID)
		t.Logf("\t\tEpisode: %s\n", ep.Episode)
		t.Logf("\t\tNumber: %d\n", ep.Number)
		t.Logf("\t\tTitle: %s\n", ep.Title)
		t.Logf("\t\tAirDate: %s\n", ep.AirDate)
		t.Logf("\t\tType: %s\n", ep.Type)
	}

	require.Equal(t, 4, len(data.Themes))
}
