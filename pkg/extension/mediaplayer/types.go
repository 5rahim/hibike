package mediaplayer

import "github.com/5rahim/hibike/pkg/extension"

type (
	// MediaPlayer is the interface that wraps the basic media player methods.
	MediaPlayer interface {
		// Start is called before the media player is used.
		Start() error
		// Play should start playing the media from the given path.
		Play(path string) error
		// Stream should start streaming the media from the given URL.
		Stream(url string) error
		// GetPlaybackStatus should return the current playback status when called.
		// It should return an error if the playback status could not be retrieved, this will cancel progress tracking.
		GetPlaybackStatus() (PlaybackStatus, error)
		// Stop will be called when the progress tracking context is canceled.
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
