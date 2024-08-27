package download

type (
	AnimeProvider interface {
		// InitUserConfig is called when the media player extension is initialized.
		// It should be used to set the user configuration.
		InitUserConfig(settings map[string]interface{})
		// Search should return the search results for the given query.
		Search(query string) ([]*AnimeItem, error)
		// GetItem should return the item details for the given item ID.
		GetItem(id string) (*AnimeItem, error)
		// Download should start downloading the item
		Download(item AnimeItem, headers map[string]string) error
	}

	AnimeItem struct {
		// ID is the unique identifier for the item.
		ID string `json:"id"`
		// Title is the title of the item.
		Title string `json:"title"`
		// URL is the URL of the item. (Optional)
		URL string `json:"url,omitempty"`
		// Image is the URL of the item image. (Optional)
		Image string `json:"image,omitempty"`
		// Description is the description of the item. (Optional)
		Description string `json:"description,omitempty"`
		// Size is the size of the item. (Optional)
		Size string `json:"size,omitempty"`
		// Resolution of the video. (Optional)
		// e.g., "1080p", "720p", "480p"
		Quality string `json:"quality,omitempty"`
		// Language is the language of the item. (Optional)
		Language string `json:"language,omitempty"`
		// Subtitle is the subtitle of the item. (Optional)
		Subtitle string `json:"subtitle,omitempty"`
		// Duration is the duration of the item. (Optional)
		Duration string `json:"duration,omitempty"`
		// UploadedAt is the date when the item was uploaded. (Optional)
		UploadedAt string `json:"uploadedAt,omitempty"`
		// DownloadURL is the URL to download the item. (Optional)
		DownloadURL string `json:"downloadURL,omitempty"`
		// DownloadHeaders are the headers for the download request. (Optional)
		DownloadHeaders map[string]string `json:"downloadHeaders,omitempty"`
	}
)
