package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/agstrc/yadb/internal/model"
	"github.com/agstrc/yadb/internal/util"
	"github.com/bwmarrin/discordgo"
)

func GetRadom() []*discordgo.MessageEmbed {

	//TODO: Handler Errors
	var response model.ModelGetRandom

	//Simple Get method to the follow url
	resp, err := http.Get(util.BaseUrl + "/manga/random")

	if err != nil {
		fmt.Println("This error has occured when requested a random manga " + "\nErr= " + err.Error())
	}

	//The Body has to be closed because of leak of conection
	// https://golang.org/pkg/net/http/#Client
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("This error has occured while reading the response " + "\nErr= " + err.Error())
	}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("This error has occured while parsing the response " + "\nErr= " + err.Error())
	}

	//fmt.Printf("%+v\n", response.Data)

	return []*discordgo.MessageEmbed{
		{
			Title:       response.Data.Attributes.Title.En,
			Description: response.Data.Attributes.Description.En,
		},
	}
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
