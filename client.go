// Package goroyale is a wrapper for the Clash Royale API at https://royaleapi.com/
package goroyale

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.royaleapi.com"

type ratelimit struct {
	remaining int
	reset     time.Time
}

// Client allows you to easily interact with RoyaleAPI.
type Client struct {
	Token string

	client    http.Client
	ratelimit ratelimit
}

// New creates a new RoyaleAPI client.
func New(token string, timeout time.Duration) (c *Client, err error) {
	c = &Client{}
	if token == "" {
		err = errors.New("client requires token for authorization with the API")
		return
	}
	c.Token = token
	if timeout == 0 {
		c.client = http.Client{Timeout: (10 * time.Second)}
	} else {
		c.client = http.Client{Timeout: (timeout)}
	}
	return
}

func (c *Client) checkRatelimit() error {
	if c.ratelimit.remaining == 0 && c.ratelimit.reset.IsZero() {
		return nil
	}
	if wait := time.Until(c.ratelimit.reset); c.ratelimit.remaining == 0 && wait > 0 {
		return newRatelimitError(wait)
	}
	return nil
}

func (c *Client) updateRatelimit(resp *http.Response) error {
	remaining := resp.Header.Get("x-ratelimit-remaining")
	if remaining != "" {
		remainingI, err := strconv.Atoi(remaining)
		if err != nil {
			return err
		}
		c.ratelimit.remaining = remainingI
	}
	retry := resp.Header.Get("x-ratelimit-retry-after")
	if retry != "" {
		sec, err := strconv.ParseInt(retry, 10, 64)
		if err != nil {
			return err
		}
		c.ratelimit.reset = time.Now().Add(time.Second * time.Duration(sec))
	}
	return nil
}

func (c *Client) get(path string, params url.Values) (bytes []byte, err error) {
	err = c.checkRatelimit()
	if err != nil {
		return
	}

	path = baseURL + path
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	req.Header.Add("auth", c.Token)
	req.URL.RawQuery = params.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var apiErr APIError
		json.Unmarshal(bytes, &apiErr)
		return []byte{}, apiErr
	}

	err = c.updateRatelimit(resp)
	return
}
