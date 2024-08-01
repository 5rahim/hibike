package autodownloader

import "github.com/5rahim/hibike/pkg/extension/torrent"

type (
	// Hook is an extension of the AutoDownloader
	Hook interface {
		// GetLatest allows the extension to return the latest torrents based on the rule
		GetLatest(rule Rule) ([]*torrent.AnimeTorrent, bool)
	}

	Rule struct {
		ComparisonTitle string   `json:"comparisonTitle"`
		ReleaseGroups   []string `json:"releaseGroups"`
		Resolutions     []string `json:"resolutions"`
	}
)
