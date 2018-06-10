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

// TODO: Add Client.GetConstants for /constants

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
// https://docs.royaleapi.com/#/endpoints/player?id=multiple-players
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

// GetPlayersBattles works like GetPlayerBattles but can return battles from multiple players.
// https://docs.royaleapi.com/#/endpoints/player_battles?id=multiple-tags
func (c Client) GetPlayersBattles(tags []string, args Args) (bs [][]Battle, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",") + "/battles"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &bs)
	}
	return
}

// GetPlayerChests gets a player's upcoming chests.
// https://docs.royaleapi.com/#/endpoints/player_chests
func (c Client) GetPlayerChests(tag string, args Args) (c PlayerChests, err error) {
	var b []byte
	path := "/player/" + tag + "/chests"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &c)
	}
	return
}

// GetPlayersChests works like GetPlayerChests but can return chests for multiple players.
// https://docs.royaleapi.com/#/endpoints/player_chests?id=multiple-players
func (c Client) GetPlayersChests(tags []string, args Args) (cs []PlayerChests, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",") + "/chests"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &cs)
	}
	return
}

// TODO: ClanSearch (https://docs.royaleapi.com/#/endpoints/clan_search)

// GetClan returns info about a specific clan.
// https://docs.royaleapi.com/#/endpoints/clan
func (c Client) GetClan(tag string, args Args) (c Clan, err error) {
	var b []byte
	path := "/clan/" + tag
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &c)
	}
	return
}

// GetClans works like GetClan but can return multiple clans.
// https://docs.royaleapi.com/#/endpoints/clan?id=multiple-clans
func (c Client) GetClans(tags []string, args Args) (cs []Clan, err error) {
	var b []byte
	path := "/clan/" + strings.Join(tags, ",")
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &cs)
	}
	return
}

// GetClanBattles returns battles played by people in the specified clan.
// https://docs.royaleapi.com/#/endpoints/clan_battles
func (c Client) GetClanBattles(tag string, args Args) (bs []Battle, err error) {
	var b []byte
	path := "/clan/" + tag + "/battles"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &bs)
	}
	return
}

// GetClanWar returns data about the current clan war.
// https://docs.royaleapi.com/#/endpoints/clan_war
func (c Client) GetClanWar(tag string, args Args) (w ClanWar, err error) {
	var b []byte
	path := "/clan/" + tag + "/war"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &w)
	}
	return
}

// GetClanWarLog returns data about past clan wars.
// https://docs.royaleapi.com/#/endpoints/clan_warlog
func (c Client) GetClanWarLog(tag string, args Args) (ws []ClanWarLogEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/warlog"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &ws)
	}
	return
}

// GetClanHistory returns a time series of member stats.
// This will only work with clans that have enabled stat tracking.
// https://docs.royaleapi.com/#/endpoints/clan_history
func (c Client) GetClanHistory(tag string, args Args) (h []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &h)
	}
	return
}

// GetClanWeeklyHistory works like GetClanHistory but returns weekly stats.
// https://docs.royaleapi.com/#/endpoints/clan_history_weekly
func (c Client) GetClanWeeklyHistory(tag string, args Args) (h []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history/weekly"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &h)
	}
	return
}

// GetClanTracking returns basic data on whether a clan is tracked.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
func (c Client) GetClanTracking(tag string, args Args) (t ClanTracking, err error) {
	var b []byte
	path := "/clan/" + tag + "/tracking"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &t)
	}
	return
}

// GetOpenTournaments returns a slice of open tournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments_open
func (c Client) GetOpenTournaments(args Args) (ts []OpenTournament, err error) {
	var b []byte
	path := "/tournaments/open"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &ts)
	}
	return
}

// GetKnownTournaments returns a slice of tournaments people have searched for.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
func (c Client) GetKnownTournaments(args Args) (ts []KnownTournament, err error) {
	var b []byte
	path := "/tournaments/known"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &ts)
	}
	return
}

// TODO: TournamentsSearch (https://docs.royaleapi.com/#/endpoints/tournaments_search)

// GetTournament returns the specified Tournament by tag.
// https://docs.royaleapi.com/#/endpoints/tournaments
func (c Client) GetTournament(tag string, args Args) (t Tournament, err error) {
	var b []byte
	path := "/tournaments/" + tag
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &t)
	}
	return
}

// GetTournaments works like GetTournament but can return multiple Tournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments?id=multiple-tournaments
func (c Client) GetTournaments(tags []string, args Args) (ts []Tournament, err error) {
	var b []byte
	path := "/tournaments/" + strings.Join(tags, ",")
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &ts)
	}
	return
}

// GetTopClans returns the top 200 clans of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_clans
func (c Client) GetTopClans(location string, args Args) (tcs []TopClan, err error) {
	var b []byte
	path := "/top/clans/" + location
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &tcs)
	}
	return
}

// GetTopPlayers returns the top 200 players of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_players
func (c Client) GetTopPlayers(location string, args Args) (tps []TopPlayer, err error) {
	var b []byte
	path := "/top/players/" + location
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &tps)
	}
	return
}

// GetPopularClans returns stats on how often a clan's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_clans
func (c Client) GetPopularClans(args Args) (pcs []PopularClan, err error) {
	var b []byte
	path := "/popular/clans"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &pcs)
	}
	return
}

// GetPopularPlayers returns stats on how often a player's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_players
func (c Client) GetPopularPlayers(args Args) (pps []PopularPlayer, err error) {
	var b []byte
	path := "/popular/players"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &pps)
	}
	return
}

// GetPopularTournaments returns stats on how often a tournament's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_tournaments
func (c Client) GetPopularTournaments(args Args) (pts []PopularTournament, err error) {
	var b []byte
	path := "/popular/tournaments"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &pts)
	}
	return
}

// GetPopularDecks returns stats on how often a deck's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_decks
func (c Client) GetPopularDecks(args Args) (pds []PopularDeck, err error) {
	var b []byte
	path := "/popular/decks"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &pds)
	}
	return
}

// GetAPIKeyStats returns information about the currently authenticated token.
// https://docs.royaleapi.com/#/endpoints/auth_stats
func (c Client) GetAPIKeyStats(args Args) (aks APIKeyStats, err error) {
	var b []byte
	path := "/auth/stats"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &aks)
	}
	return
}

// GetEndpoints returns all the available endpoints for the API.
// It does not have any special incorporation with this wrapper and is simply included for completion's sake.
// https://docs.royaleapi.com/#/endpoints/endpoints
func (c Client) GetEndpoints(args Args) (eps []string, err error) {
	var b []byte
	path := "/endpoints"
	if b, err = c.get(path, args); err == nil {
		err = json.Unmarshal(b, &eps)
	}
	return
}
