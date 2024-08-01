package mediaplayer

type (
	// MediaPlayer is the interface that wraps the basic media player methods.
	MediaPlayer interface {
		// InitConfig is called when the media player extension is initialized.
		// It should be used to set the user configuration.
		InitConfig(config map[string]interface{})
		GetSettings() Settings
		// Start is called before the media player is used.
		Start() error
		// Play should start playing the media from the given path.
		Play(req PlayRequest) (PlayResponse, error)
		// Stream should start streaming the media from the given URL.
		Stream(req PlayRequest) (PlayResponse, error)
		// GetPlaybackStatus should return the current playback status when called.
		// It should return an error if the playback status could not be retrieved, this will cancel progress tracking.
		GetPlaybackStatus() (PlaybackStatus, error)
	}

	ClientInfo struct {
		UserAgent string `json:"userAgent"`
		// Platform is the operating system of the client.
		// It "will be one of the following: "windows", "mac", "linux", "android", "ios", "tvos", "webos", "tizen"
		Platform string `json:"platform"`
	}

	PlayRequest struct {
		// URL or file path of the media.
		Path       string     `json:"path"`
		ClientInfo ClientInfo `json:"clientInfo"`
	}

	// PlayResponse is the response returned by the Play and Stream methods.
	// It contains the command to be executed or the URL to be opened.
	PlayResponse struct {
		// Command to be executed. (Optional)
		// This requires "exec" permission.
		Cmd string `json:"cmd,omitempty"`
		// URL to be opened. (Optional)
		// This is used if the media player is a mobile app.
		OpenURL string `json:"openURL,omitempty"`
	}

	Settings struct {
		// If true, GetPlaybackStatus should return the current playback status when called.
		// If false, the user will be prompted to manually track the progress.
		CanTrackProgress bool `json:"canTrackProgress"`
	}

	PlaybackStatus struct {
		// Completion percentage of the media.
		// It should be a float from 0 to 100.
		CompletionPercentage float64 `json:"completionPercentage"`
		// Whether the media is currently playing.
		Playing bool `json:"playing"`
		// Current media filename.
		Filename string `json:"filename"`
		// Current media path.
		Path string `json:"path"`
		// Duration of the media in milliseconds.
		Duration int `json:"duration"`
		// Current media absolute file path.
		Filepath string `json:"filepath"`
	}
)
