package goroyale

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const BASE_URL = "https://api.royaleapi.com"

// Client allows you to easily interact with RoyaleAPI.
type Client struct {
	Token string

	client http.Client
}

// New creates a new RoyaleAPI client.
func New(token string) (c *Client, err error) {
	if token == "" {
		panic("Client requires token for authorization with the API.")
	}
	c.Token = token
	c.client = http.Client{Timeout: (10 * time.Second)}

	return
}

func (c Client) get(endpoint string, args Args) (resp *http.Response, err error) {
	path := BASE_URL + endpoint
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}

	req.Header.Add("auth", c.Token)
	resp, err = c.client.Do(req)
	return
}

// Args represents special args to pass in the request.
// The API supports args for Field Filter https://docs.royaleapi.com/#/field_filter
// and Pagination https://docs.royaleapi.com/#/pagination.
type Args struct {
	Keys    []string
	Exclude []string
	Max     int
	Page    int
}

// GetAPIVersion requests the current version of the API.
// https://docs.royaleapi.com/#/endpoints/version
func (c Client) GetAPIVersion() (bodyString string, err error) {
	if args != Args{} {
		log.
	}	

	resp, err := c.get("/version", Args{})
	defer resp.Body.Close()
	if err != nil {
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString = string(bodyBytes)
	return
}

// GetPlayer retrieves a player by their tag.
// https://docs.royaleapi.com/#/endpoints/player
func (c Client) GetPlayer(tag string, args Args) (p Player, err error) {
	path := "/player/" + tag
	resp, err := c.get(path, args)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &p)
	return
}
