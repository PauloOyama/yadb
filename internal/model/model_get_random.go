package model

import "time"

type ModelGetRandom struct {
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
