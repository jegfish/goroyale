package goroyale

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type params interface {
	toQuery() url.Values
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

func (args Args) toQuery() url.Values {
	q := url.Values{}
	if args.Keys != nil {
		q.Set("keys", strings.Join(args.Keys, ","))
	}

	if args.Exclude != nil {
		q.Set("exclude", strings.Join(args.Keys, ","))
	}

	if args.Max != 0 {
		q.Set("max", string(args.Max))
	}

	if args.Page != 0 {
		q.Set("page", string(args.Page))
	}

	return q
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

func (c *Client) get(path string, args params) (bytes []byte, err error) {
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
	req.URL.RawQuery = args.toQuery().Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	err = c.updateRatelimit(resp)
	return
}

// ClanSearchArgs lets you provide search terms for ClanSearch.
// https://docs.royaleapi.com/#/endpoints/clan_search
type ClanSearchArgs struct {
	Args

	Name       string
	MinScore   int
	MinMembers int
	MaxMembers int
	LocationID int
}

func (args ClanSearchArgs) toQuery() url.Values {
	q := url.Values{}
	if args.Keys != nil {
		q.Set("keys", strings.Join(args.Keys, ","))
	}

	if args.Exclude != nil {
		q.Set("exclude", strings.Join(args.Keys, ","))
	}

	if args.Max != 0 {
		q.Set("max", string(args.Max))
	}

	if args.Page != 0 {
		q.Set("page", string(args.Page))
	}

	if args.Name != "" {
		q.Set("name", args.Name)
	}

	if args.MinScore != 0 {
		q.Set("score", string(args.MinScore))
	}

	if args.MinMembers != 0 {
		q.Set("minMembers", string(args.MinMembers))
	}

	if args.MaxMembers != 0 {
		q.Set("maxMembers", string(args.MaxMembers))
	}

	if args.LocationID != 0 {
		q.Set("locationId", string(args.LocationID))
	}

	return q
}

// TournamentSearchArgs lets you provide search terms for TournamentSearch
// https://docs.royaleapi.com/#/endpoints/tournaments_search
type TournamentSearchArgs struct {
	Args

	Name string
}

func (args TournamentSearchArgs) toQuery() url.Values {
	q := url.Values{}
	if args.Keys != nil {
		q.Set("keys", strings.Join(args.Keys, ","))
	}

	if args.Exclude != nil {
		q.Set("exclude", strings.Join(args.Keys, ","))
	}

	if args.Max != 0 {
		q.Set("max", string(args.Max))
	}

	if args.Page != 0 {
		q.Set("page", string(args.Page))
	}

	if args.Name != "" {
		q.Set("name", args.Name)
	}

	return q
}
