package goroyale

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

// Args represents special args to pass in the request.
// The API supports args for Field Filter https://docs.royaleapi.com/#/field_filter
// and Pagination https://docs.royaleapi.com/#/pagination.
type Args struct {
	Keys    []string
	Exclude []string
	Max     int
	Page    int
}

func argQuery(args Args) (q url.Values) {
	if args.Keys != nil {
		q.Add("keys", strings.Join(args.Keys, ","))
	}

	if args.Exclude != nil {
		q.Add("exclude", strings.Join(args.Keys, ","))
	}

	if args.Max != 0 {
		q.Add("max", string(args.Max))
	}

	if args.Page != 0 {
		q.Add("page", string(args.Page))
	}

	return
}

func (c Client) get(path string, args Args) (bytes []byte, err error) {
	path = BASE_URL + path
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	req.Header.Add("auth", c.Token)
	req.URL.RawQuery = argQuery(args).Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	return
}

// GetAPIVersion requests the current version of the API.
// https://docs.royaleapi.com/#/endpoints/version
func (c Client) GetAPIVersion() (bodyString string, err error) {
	bytes, err := c.get("/version", Args{})
	if err != nil {
		return
	}
	bodyString = string(bytes)
	return
}

// GetPlayer retrieves a player by their tag.
// https://docs.royaleapi.com/#/endpoints/player
func (c Client) GetPlayer(tag string, args Args) (p Player, err error) {
	var b []byte
	path := "/player/" + tag
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &p)
	}
	return
}

// GetPlayers works like GetPlayer but can return multiple players.
// The API asks that you don't include more than 7 tags in this request.
// https://docs.royaleapi.com/#/endpoints/player
func (c Client) GetPlayers(tags []string, args Args) (ps []Player, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",")
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &ps)
	}
	return
}

// GetPlayerBattles gets battles a player participated in.
// https://docs.royaleapi.com/#/endpoints/player_battles
func (c Client) GetPlayerBattles(tag string, args Args) (bs []Battle, err error) {
	var b []byte
	path := "/player/" + tag + "/battles"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &bs)
	}
	return
}
