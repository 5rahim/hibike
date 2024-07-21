package manga

type (
	Provider interface {
		Search(opts SearchOptions) ([]*SearchResult, error)
		FindChapters(id string) ([]*ChapterDetails, error)
		FindChapterPages(id string) ([]*ChapterPage, error)
	}

	SearchOptions struct {
		Query string
		Year  int
	}

	SearchResult struct {
		ID           string
		Title        string
		Synonyms     []string
		Year         int
		Image        string
		Provider     string
		SearchRating float64
	}

	ChapterDetails struct {
		Provider  string `json:"provider"`
		ID        string `json:"id"`
		URL       string `json:"url"`
		Title     string `json:"title"`
		Chapter   string `json:"chapter"` // e.g., "1", "1.5", "2", "3"
		Index     uint   `json:"index"`   // Index of the chapter in the manga
		Rating    int    `json:"rating"`
		UpdatedAt string `json:"updatedAt"`
	}

	ChapterPage struct {
		Provider string            `json:"provider"`
		URL      string            `json:"url"`
		Index    int               `json:"index"`
		Headers  map[string]string `json:"headers"`
	}
)
