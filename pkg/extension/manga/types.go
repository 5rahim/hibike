package manga

import "github.com/5rahim/hibike/pkg/extension"

type (
	Provider interface {
		// Search returns the search results for the given query.
		Search(opts SearchOptions) ([]*SearchResult, error)
		// FindChapters returns the chapter details for the given manga ID.
		FindChapters(id string) ([]*ChapterDetails, error)
		// FindChapterPages returns the chapter pages for the given chapter ID.
		FindChapterPages(id string) ([]*ChapterPage, error)
		// GetSettings returns the provider settings.
		GetSettings() Settings
	}

	Settings struct {
		SupportsMultiGroup    bool `json:"supportsMultiGroup"`
		SupportsMultiLanguage bool `json:"supportsMultiLanguage"`
		// Groups (or Scanlators) supported by the extension.
		// The one selected by the user will be passed to [Provider.Search].
		// The default group should be the first option.
		Groups []extension.SelectOption `json:"groups,omitempty"`
		// Languages supported by the extension.
		// The one selected by the user will be passed to [Provider.Search].
		// Leave it empty if the extension does not support languages.
		// The default language should be the first option.
		Languages []extension.SelectOption `json:"languages,omitempty"`
	}

	SearchOptions struct {
		Query string `json:"query"`
		// Year is the year the manga was released.
		// It will be 0 if the year is not available.
		Year int `json:"year"`
		// Language selected by the user.
		Language string `json:"language"`
		// Group selected by the user.
		Group string `json:"group"`
	}

	SearchResult struct {
		// "ID" of the extension.
		Provider string `json:"provider"`
		// Language of the manga.
		// Leave it empty if the language is not available.
		Language string `json:"language,omitempty"`
		// It is used to fetch the chapter details.
		// It can be a combination of keys separated by a delimiter. (Delimiters should not be slashes).
		//	If the extension supports multiple languages, the language key should be included. (e.g., "one-piece$en").
		//	If the extension supports multiple groups, the group key should be included. (e.g., "one-piece$group-1").
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

	ChapterDetails struct {
		// "ID" of the extension.
		// This should be the same as the extension ID and follow the same format.
		Provider string `json:"provider"`
		// ID of the chapter, used to fetch the chapter pages.
		// It can be a combination of keys separated by a delimiter. (Delimiters should not be slashes).
		//	If the extension supports multiple languages, the language key should be included. (e.g., "one-piece-001$chapter-1$en").
		//	If the extension supports multiple groups, the group key should be included. (e.g., "one-piece-001$chapter-1$group-1").
		ID string `json:"id"`
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
