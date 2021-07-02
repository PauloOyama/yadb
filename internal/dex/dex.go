// package dex implements integration functions with the MangaDex API
package dex

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/agstrc/yadb/internal/util"
	"github.com/bwmarrin/discordgo"
)

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

//getFileName represents the JSON return by MangaDex on it's `/cover/{cover_art_id}` endpoint
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

//getAutorName represents the JSON return by MangaDex on it's `/author/{author_id}`endpoint
type getAutorName struct {
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

//getSearchManga represents the JSON return by MangaDex on it's `/manga/{mangas_id}` endpoint
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

//getAuthor receives the author's Id and return the his name using the MangaDex API
func getAuthor(authorId string) (string, error) {

	var response getAutorName

	resp, err := http.Get(util.BaseUrl + "/author/" + authorId)

	if err != nil {
		return "Autor Not Found", fmt.Errorf("failed to send a get request to the upstream API: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error EOF", fmt.Errorf("failed to read response body returned by upstream: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "Error when Parsing the response", fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	return response.Data.Attributes.Name, nil

}

func getCoverImage(coverId string) (string, error) {

	var response getFileName

	resp, err := http.Get(util.BaseUrl + "/cover/" + coverId)

	if err != nil {
		return "Image Not Found", fmt.Errorf("failed to send a get request to the upstream API: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error EOF", fmt.Errorf("failed to read response body returned by upstream: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "Error when Parsing the response", fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	return response.Data.Attributes.FileName, nil
}

// GetRandom returns a slice of *discordgo.MessageEmbed with information about a random Manga acquired from the
// MangaDex API
func GetRandom() ([]*discordgo.MessageEmbed, error) {
	var response RandomManga
	var autorID string
	var coverID string

	resp, err := http.Get(util.BaseUrl + "/manga/random")

	if err != nil {
		return nil, fmt.Errorf("failed to send a get request to the upstream API: %w", err)
	}

	// The Body has to be closed to avoid a memory leak
	// https://golang.org/pkg/net/http/#Client
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body returned by upstream: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	for _, s := range response.Relationships {
		if s.Type == "author" {
			autorID = s.ID
		}
	}

	autorName, err := getAuthor(autorID)

	if err != nil {
		autorName = ""
	}

	for _, s := range response.Relationships {
		if s.Type == "cover_art" {
			coverID = s.ID
		}
	}

	fileName, err := getCoverImage(coverID)

	//Handle error
	if err != nil {
		fileName = ""
	}

	fmt.Println("https://uploads.mangadex.org/covers/" + response.Data.ID + "/" + fileName + ".512.jpg")
	return []*discordgo.MessageEmbed{
		{
			URL:         "https://mangadex.org/title/" + response.Data.ID,
			Title:       response.Data.Attributes.Title.En,
			Description: response.Data.Attributes.Description.En,
			Image: &discordgo.MessageEmbedImage{
				URL: "https://uploads.mangadex.org/covers/" + response.Data.ID + "/" + fileName + ".512.jpg",
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: autorName,
			},
		},
	}, nil
}

//GetMangaReader receives the string received from the slashCommand ´get-manga´ and search this title 
//using MangaDex API, this function return a []*discordgo.MessageEmbed with the response value and an error.
//Important: If the title searched is not found the API will return an empty response 
func GetMangaReader(title string) ([]*discordgo.MessageEmbed, error) {

	var client http.Client
	var response getSearchManga
	var autorID string
	var coverID string

	//Make a simple request with the manga to be found
	req, err := http.NewRequest("GET", util.BaseUrl+"/manga", nil)

	if err != nil {
		return nil, fmt.Errorf("Error while trying to build a request"+"\nErr=%w", err)
	}

	// Build parameter
	params := url.Values{}
	params.Add("title", title)

	req.URL.RawQuery = params.Encode()

	//Send de request
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Error when sent a request do found the manga " + title + "\nErr= " + err.Error())
	}

	//Parse the receveid Json to be readed
	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("This error has occured while reading the response " + "\nErr= " + err.Error())
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	if len(response.Results) == 0 {
		return []*discordgo.MessageEmbed{
			{
				Title:       "Manga não encontrado",
				Description: "O Mangá solicitado não pode ser encontrado, verifique se o nome do mangá está correto!",
				Image: &discordgo.MessageEmbedImage{
					URL: "https://blog.golang.org/5years/gophers5th.jpg",
				},
			},
		}, nil
	}

	//For now, it only return the first value of the list in the response.Results, it seems that the function 
	//append causes some misunderstanding in the return function

	for _, s := range response.Results[0].Relationships {
		if s.Type == "author" {
			autorID = s.ID
		}
	}

	autorName, err := getAuthor(autorID)

	if err != nil {
		autorName = ""
	}

	for _, s := range response.Results[0].Relationships {
		if s.Type == "cover_art" {
			coverID = s.ID
		}
	}

	fileName, err := getCoverImage(coverID)

	//Handle error
	if err != nil {
		fileName = ""
	}

	result := []*discordgo.MessageEmbed{
		{

			Title:       response.Results[0].Data.Attributes.Title.En,
			Description: response.Results[0].Data.Attributes.Description.En,
			Image: &discordgo.MessageEmbedImage{
				URL: "https://uploads.mangadex.org/covers/" + response.Results[0].Data.ID + "/" + fileName + ".512.jpg",
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: autorName,
			},
		},
	}

	// for _, s := range response.Results {

	// 	fmt.Println(s.Data.Attributes.Title.En)

	// 	part := &discordgo.MessageEmbed{
	// 		Title:       s.Data.Attributes.Title.En,
	// 		Description: s.Data.Attributes.Description.En,
	// 		Image: &discordgo.MessageEmbedImage{
	// 			URL: "https://uploads.mangadex.org/covers/" + response.Results[0].Data.ID + "/" + fileName + ".512.jpg",
	// 		},
	// 		Author: &discordgo.MessageEmbedAuthor{
	// 			Name: autorName,
	// 		},
	// 	}
	// 	result = append(result, part)
	// }

	return result, nil
}
