package manga

const (
	ChapterFilterLanguage ChapterFilter = "language"
	ChapterFilterGroup    ChapterFilter = "group"
)

type (
	ChapterFilter string

	SelectOption struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	Provider interface {
		// Search returns the search results for the given query.
		Search(opts SearchOptions) ([]*SearchResult, error)
		// FindChapters returns the chapter details for the given manga ID.
		FindChapters(opts FindChapterOptions) ([]*ChapterDetails, error)
		// FindChapterPages returns the chapter pages for the given chapter ID.
		FindChapterPages(id string) ([]*ChapterPage, error)
		// GetSettings returns the provider settings.
		GetSettings() Settings
	}

	Settings struct {
		ChapterFilters []ChapterFilter `json:"chapterFilters"`
		Languages      []SelectOption  `json:"languages,omitempty"`
		Groups         []SelectOption  `json:"groups,omitempty"`
	}

	SearchOptions struct {
		Query string `json:"query"`
		// Year is the year the manga was released.
		// It will be 0 if the year is not available.
		Year int `json:"year"`
		// Language requested by the user.
		// It will be empty if the language is not available.
		Language string `json:"language,omitempty"`
		// Group requested by the user.
		// It will be empty if the group is not available.
		Group string `json:"group,omitempty"`
	}

	SearchResult struct {
		// "ID" of the extension.
		Provider string `json:"provider"`
		// Language of the manga.
		// Leave it empty if the language is not available.
		Language string `json:"language,omitempty"`
		// It is used to fetch the chapter details.
		// It can be a combination of keys separated by the $ delimiter.
		ID string `json:"id"`
		// The title of the manga.
		Title string `json:"title"`
		// Synonyms are alternative titles for the manga.
		Synonyms []string `json:"synonyms,omitempty"`
		// Year the manga was released.
		Year int `json:"year,omitempty"`
		// URL of the manga cover image.
		Image string `json:"image,omitempty"`
		// Indicates how well the chapter title matches the search query.
		// It is a number from 0 to 1.
		// Leave it empty if the comparison should be done by Seanime.
		SearchRating float64 `json:"searchRating,omitempty"`
	}

	FindChapterOptions struct {
		// ID is the manga slug.
		ID string `json:"id"`
		// Language of the manga.
		// Will be empty if the language is not available.
		Language string `json:"language,omitempty"`
		// Group of the manga.
		// Will be empty if the group is not available.
		Group string `json:"group,omitempty"`
	}

	ChapterDetails struct {
		// ID of the provider.
		// This should be the same as the extension ID and follow the same format.
		Provider string `json:"provider"`
		// ID is the chapter slug.
		// It is used to fetch the chapter pages.
		// It can be a combination of keys separated by the $ delimiter.
		// e.g., "10010$one-piece-1", where "10010" is the manga ID and "one-piece-1" is the chapter slug that is reconstructed to "%url/10010/one-piece-1".
		ID string `json:"id"`
		// Language of the chapter.
		// Leave it empty if the language is not available.
		Language string `json:"language,omitempty"`
		// The chapter page URL.
		URL string `json:"url"`
		// The chapter title.
		// It should be in this format: "Chapter X.Y - {title}" where X is the chapter number and Y is the subchapter number.
		Title string `json:"title"`
		// e.g., "1", "1.5", "2", "3"
		Chapter string `json:"chapter"`
		// From 0 to n
		Index uint `json:"index"`
		// The rating of the chapter. It is a number from 0 to 100.
		// Leave it empty if the rating is not available.
		Rating int `json:"rating,omitempty"`
		// UpdatedAt is the date when the chapter was last updated.
		// It should be in the format "YYYY-MM-DD".
		// Leave it empty if the date is not available.
		UpdatedAt string `json:"updatedAt,omitempty"`
	}

	ChapterPage struct {
		// ID of the provider.
		// This should be the same as the extension ID and follow the same format.
		Provider string `json:"provider"`
		// URL of the chapter page.
		URL string `json:"url"`
		// Index of the page in the chapter.
		// From 0 to n.
		Index int `json:"index"`
		// Request headers for the page if proxying is required.
		Headers map[string]string `json:"headers"`
	}
)
