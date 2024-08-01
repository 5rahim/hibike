package mediaplayer

import "github.com/5rahim/hibike/pkg/extension"

type (
	// MediaPlayer is the interface that wraps the basic media player methods.
	MediaPlayer interface {
		Start() error
		Play(path string) error
		Stream(path string) error
		GetPlaybackStatus() (PlaybackStatus, error)
		Stop() error
		GetSettings() Settings
		GetSettingInputs() []extension.SettingInput
	}

	Settings struct {
		// If true, GetPlaybackStatus should return the current playback status when called.
		CanTrackProgress bool `json:"canTrackProgress"`
	}

	PlaybackStatus struct {
		CompletionPercentage float64 `json:"completionPercentage"`
		Playing              bool    `json:"playing"`
		Filename             string  `json:"filename"`
		Path                 string  `json:"path"`
		Duration             int     `json:"duration"` // in ms
		Filepath             string  `json:"filepath"`
	}
)
