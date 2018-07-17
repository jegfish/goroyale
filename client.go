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

// Client allows you to easily interact with RoyaleAPI.
type Client struct {
	Token string

	client http.Client
	// using empty struct because it has a byte size of 0
	// i don't care what's in the channel, just that something is
	rateBucket chan struct{}
}

// New creates a new RoyaleAPI client.
func New(token string, timeout time.Duration) (c *Client, err error) {
	c = &Client{
		client:     http.Client{Timeout: 10 * time.Second},
		rateBucket: make(chan struct{}, 5),
	}
	if token == "" {
		err = errors.New("client requires token for authorization with the API")
		return
	}
	c.Token = token
	if timeout != 0 {
		c.client = http.Client{Timeout: timeout}
	}

	// Allow initial request
	c.rateBucket <- struct{}{}
	return
}

func (c *Client) updateRatelimit(resp *http.Response) error {
	remaining := resp.Header.Get("x-ratelimit-remaining")
	if remaining != "" {
		remainingI, err := strconv.Atoi(remaining)
		if err != nil {
			return err
		}

		if remainingI > 0 {
			c.rateBucket <- struct{}{}
		}
	}
	retry := resp.Header.Get("x-ratelimit-retry-after")
	if retry != "" {
		sec, err := strconv.ParseInt(retry, 10, 64)
		if err != nil {
			return err
		}

		// Ratelimit-Retry-After only shows up when Ratelimit-Remaining hits 0
		// Wait until next request is available and add it to the rateBucket
		go func() {
			time.Sleep(time.Duration(sec) * time.Second)
			c.rateBucket <- struct{}{}
		}()
	}
	return nil
}

func (c *Client) get(path string, params url.Values) (bytes []byte, err error) {
	// take one request out of the rateBucket
	<-c.rateBucket

	path = baseURL + path
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	req.Header.Add("auth", c.Token)
	req.URL.RawQuery = params.Encode()

	resp, err := c.client.Do(req)
	defer c.updateRatelimit(resp)
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

	return
}
