package goroyale

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.royaleapi.com"

type ratelimit struct {
	remaining int
	reset     int64
}

// Client allows you to easily interact with RoyaleAPI.
type Client struct {
	Token string

	client    http.Client
	ratelimit ratelimit
}

// New creates a new RoyaleAPI client.
func New(token string, timeout time.Duration) (c Client, err error) {
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
	if c.ratelimit.remaining == 0 || c.ratelimit.reset == 0 {
		return nil
	}
	if now := time.Now().UnixNano() / 1000000; c.ratelimit.remaining == 0 && now < c.ratelimit.reset {
		return fmt.Errorf("ratelimit, retry in: %d", c.ratelimit.reset-now)
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
	reset := resp.Header.Get("x-ratelimit-reset")
	if reset != "" {
		resetI, err := strconv.ParseInt(reset, 10, 64)
		if err != nil {
			return err
		}
		c.ratelimit.reset = resetI
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
	err = c.updateRatelimit(resp)
	return
}
