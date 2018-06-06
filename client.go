package goroyale

import (
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

func (c Client) get(endpoint string, keys []string, exclude []string) (resp *http.Response, err error) {
	path := BASE_URL + endpoint
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}

	req.Header.Add("auth", c.Token)
	resp, err = c.client.Do(req)
	return
}
