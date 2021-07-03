package dex

import "time"

// RandomManga represents the JSON returned by MangaDex on its `/manga/random` endpoint
type RandomManga struct {
	Result string `json:"result"`
	Data   struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title struct {
				En string `json:"en"`
			} `json:"title"`
			AltTitles   []interface{} `json:"altTitles"`
			Description struct {
				En string `json:"en"`
			} `json:"description"`
			Links                  interface{} `json:"links"`
			OriginalLanguage       string      `json:"originalLanguage"`
			LastVolume             interface{} `json:"lastVolume"`
			LastChapter            interface{} `json:"lastChapter"`
			PublicationDemographic interface{} `json:"publicationDemographic"`
			Status                 string      `json:"status"`
			Year                   interface{} `json:"year"`
			ContentRating          string      `json:"contentRating"`
			Tags                   []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name struct {
						En string `json:"en"`
					} `json:"name"`
					Description []interface{} `json:"description"`
					Group       string        `json:"group"`
					Version     int           `json:"version"`
				} `json:"attributes"`
			} `json:"tags"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			Version   int       `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
	Relationships []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"relationships"`
}

// getFileName represents the JSON returned by MangaDex on its `/cover/{cover_art_id}` endpoint
type getFileName struct {
	Result string `json:"result"`
	Data   struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Description string      `json:"description"`
			Volume      interface{} `json:"volume"`
			FileName    string      `json:"fileName"`
			CreatedAt   time.Time   `json:"createdAt"`
			UpdatedAt   time.Time   `json:"updatedAt"`
			Version     int         `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
	Relationships []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"relationships"`
}

// getAuthorName represents the JSON return by MangaDex on its `/author/{author_id}`endpoint
type getAuthorName struct {
	Result string `json:"result"`
	Data   struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name      string        `json:"name"`
			ImageURL  interface{}   `json:"imageUrl"`
			Biography []interface{} `json:"biography"`
			CreatedAt time.Time     `json:"createdAt"`
			UpdatedAt time.Time     `json:"updatedAt"`
			Version   int           `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
	Relationships []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"relationships"`
}

// getSearchManga represents the JSON return by MangaDex on its `/manga/{mangas_id}` endpoint
type getSearchManga struct {
	Results []struct {
		Result string `json:"result"`
		Data   struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Attributes struct {
				Title struct {
					En string `json:"en"`
				} `json:"title"`
				AltTitles []struct {
					En string `json:"en"`
				} `json:"altTitles"`
				Description struct {
					De string `json:"de"`
					En string `json:"en"`
					Fr string `json:"fr"`
					It string `json:"it"`
					Ro string `json:"ro"`
					Ru string `json:"ru"`
					Tr string `json:"tr"`
				} `json:"description"`
				Links struct {
					Al    string `json:"al"`
					Ap    string `json:"ap"`
					Kt    string `json:"kt"`
					Mu    string `json:"mu"`
					Nu    string `json:"nu"`
					Mal   string `json:"mal"`
					Raw   string `json:"raw"`
					Engtl string `json:"engtl"`
				} `json:"links"`
				OriginalLanguage       string      `json:"originalLanguage"`
				LastVolume             interface{} `json:"lastVolume"`
				LastChapter            interface{} `json:"lastChapter"`
				PublicationDemographic string      `json:"publicationDemographic"`
				Status                 string      `json:"status"`
				Year                   interface{} `json:"year"`
				ContentRating          string      `json:"contentRating"`
				Tags                   []struct {
					ID         string `json:"id"`
					Type       string `json:"type"`
					Attributes struct {
						Name struct {
							En string `json:"en"`
						} `json:"name"`
						Description []interface{} `json:"description"`
						Group       string        `json:"group"`
						Version     int           `json:"version"`
					} `json:"attributes"`
				} `json:"tags"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
				Version   int       `json:"version"`
			} `json:"attributes"`
		} `json:"data"`
		Relationships []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"relationships"`
	} `json:"results"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
