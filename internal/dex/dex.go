// package dex implements integration functions with the MangaDex API
package dex

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

// GetRandom returns a slice of *discordgo.MessageEmbed with information about a random Manga acquired from the
// MangaDex API
func GetRandom() ([]*discordgo.MessageEmbed, error) {
	var response RandomManga

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

	return []*discordgo.MessageEmbed{
		{
			Title:       response.Data.Attributes.Title.En,
			Description: response.Data.Attributes.Description.En,
		},
	}, nil
}

func GetManga(title string) {

	var client http.Client

	// Build parameter
	params := url.Values{}
	params.Add("Title", title)

	//Make a simple request with the manga to be found
	req, err := http.NewRequest("GET", util.BaseUrl+"/manga", strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println("Error while trying to build a request"+"\nErr=", err)
	}

	//Send de request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error when sent a request do found the manga " + title + "\nErr= " + err.Error())
	}

	//Parse the receveid Json to be readed
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("This error has occured while reading the response " + "\nErr= " + err.Error())
	}

	defer resp.Body.Close()
	fmt.Println(string(body))
}
