// package dex implements integration functions with the MangaDex API
package dex

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/agstrc/yadb/internal/util"
	dg "github.com/bwmarrin/discordgo"
)

//getAuthor receives the author's Id and return the his name using the MangaDex API
func getAuthor(authorId string) (string, error) {
	var response getAuthorName

	resp, err := http.Get(util.BaseUrl + "/author/" + authorId)

	if err != nil {
		return "", fmt.Errorf("failed to send a get request to the upstream API: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body returned by upstream: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal the response body: %w", err)
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

// GetRandom returns a slice of *dg.MessageEmbed with information about a random Manga acquired from the
// MangaDex API
func GetRandom() ([]*dg.MessageEmbed, error) {
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
	var cover *dg.MessageEmbedImage

	// if no error was returned, we instantiate the cover embed. Otherwise, keep it nil
	if err == nil {
		cover = &dg.MessageEmbedImage{URL: "https://uploads.mangadex.org/covers/" +
			response.Data.ID + "/" + fileName + ".512.jpg"}
	}

	return []*dg.MessageEmbed{
		{
			URL:         "https://mangadex.org/title/" + response.Data.ID,
			Title:       response.Data.Attributes.Title.En,
			Description: response.Data.Attributes.Description.En,
			Image:       cover,
			Author: &dg.MessageEmbedAuthor{
				Name: autorName,
			},
		},
	}, nil
}

// GetMangaReader receives the string received from the slashCommand ´get-manga´ and search this title
// using MangaDex API, this function return a []*dg.MessageEmbed with the response value and an error.
// Important: If the title searched is not found the API will return an empty response
func GetMangaReader(title string) ([]*dg.MessageEmbed, error) {

	var response getSearchManga
	var authorID string
	var coverID string

	//Make a simple request with the manga to be found
	req, err := http.NewRequest("GET", util.BaseUrl+"/manga", nil)

	if err != nil {
		return nil, fmt.Errorf("error while trying to build a request: %w", err)
	}

	// Build parameter
	params := url.Values{}
	params.Add("title", title)

	req.URL.RawQuery = params.Encode()

	// Send the request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("a GET (%s) request failed: %w", req.URL.String(), err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}

	defer resp.Body.Close()

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	if len(response.Results) == 0 {
		return []*dg.MessageEmbed{
			{
				Title:       "Manga não encontrado",
				Description: "O Mangá solicitado não pode ser encontrado, verifique se o nome do mangá está correto!",
				Image: &dg.MessageEmbedImage{
					URL: "https://blog.golang.org/5years/gophers5th.jpg",
				},
			},
		}, nil
	}

	//For now, it only return the first value of the list in the response.Results, it seems that the function
	//append causes some misunderstanding in the return function

	for _, s := range response.Results[0].Relationships {
		if s.Type == "author" {
			authorID = s.ID
		}
	}

	authorName, _ := getAuthor(authorID)

	for _, s := range response.Results[0].Relationships {
		if s.Type == "cover_art" {
			coverID = s.ID
		}
	}

	fileName, err := getCoverImage(coverID)
	var cover *dg.MessageEmbedImage

	if err == nil {
		cover = &dg.MessageEmbedImage{URL: "https://uploads.mangadex.org/covers/" +
			response.Results[0].Data.ID + "/" + fileName + ".512.jpg"}
	}

	result := []*dg.MessageEmbed{
		{

			Title:       response.Results[0].Data.Attributes.Title.En,
			Description: response.Results[0].Data.Attributes.Description.En,
			Image:       cover,
			Author: &dg.MessageEmbedAuthor{
				Name: authorName,
			},
		},
	}

	for _, s := range response.Results {
		part := &dg.MessageEmbed{
			Title:       s.Data.Attributes.Title.En,
			Description: s.Data.Attributes.Description.En,
			Image: &dg.MessageEmbedImage{
				URL: "https://uploads.mangadex.org/covers/" + response.Results[0].Data.ID + "/" + fileName + ".512.jpg",
			},
			Author: &dg.MessageEmbedAuthor{
				Name: authorName,
			},
		}
		result = append(result, part)
	}

	return result, nil
}
