package goroyale

import (
	"encoding/json"
	"net/url"
	"strings"
)

// APIVersion requests the current version of the API.
// https://docs.royaleapi.com/#/endpoints/version
func (c *Client) APIVersion() (ver string, err error) {
	bytes, err := c.get("/version", url.Values{})
	if err != nil {
		return
	}
	ver = string(bytes)
	return
}

// Constants returns constants from the API.
// https://docs.royaleapi.com/#/endpoints/constants
func (c *Client) Constants(params url.Values) (constants Constants, err error) {
	var b []byte
	path := "/constants"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &constants)
	}
	return
}

// Player retrieves a player by their tag.
// https://docs.royaleapi.com/#/endpoints/player
func (c *Client) Player(tag string, params url.Values) (player Player, err error) {
	var b []byte
	path := "/player/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &player)
	}
	return
}

// Players works like Player but can return multiple players.
// The API asks that you don't include more than 7 tags in this request.
// https://docs.royaleapi.com/#/endpoints/player?id=multiple-players
func (c *Client) Players(tags []string, params url.Values) (players []Player, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &players)
	}
	return
}

// PlayerBattles s battles a player participated in.
// https://docs.royaleapi.com/#/endpoints/player_battles
func (c *Client) PlayerBattles(tag string, params url.Values) (battles []Battle, err error) {
	var b []byte
	path := "/player/" + tag + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// PlayersBattles works like PlayerBattles but can return battles from multiple players.
// https://docs.royaleapi.com/#/endpoints/player_battles?id=multiple-tags
func (c *Client) PlayersBattles(tags []string, params url.Values) (battles [][]Battle, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",") + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// PlayerChests s a player's upcoming chests.
// https://docs.royaleapi.com/#/endpoints/player_chests
func (c *Client) PlayerChests(tag string, params url.Values) (chests PlayerChests, err error) {
	var b []byte
	path := "/player/" + tag + "/chests"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &chests)
	}
	return
}

// PlayersChests works like PlayerChests but can return chests for multiple players.
// https://docs.royaleapi.com/#/endpoints/player_chests?id=multiple-players
func (c *Client) PlayersChests(tags []string, params url.Values) (chests []PlayerChests, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",") + "/chests"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &chests)
	}
	return
}

// ClanSearch searches for a clan using the provided parmameters.
// https://docs.royaleapi.com/#/endpoints/clan_search
func (c *Client) ClanSearch(params url.Values) (clans []ClanSearch, err error) {
	var b []byte
	path := "clan/search"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clans)
	}
	return
}

// Clan returns info about a specific clan.
// https://docs.royaleapi.com/#/endpoints/clan
func (c *Client) Clan(tag string, params url.Values) (clan Clan, err error) {
	var b []byte
	path := "/clan/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clan)
	}
	return
}

// Clans works like Clan but can return multiple clans.
// https://docs.royaleapi.com/#/endpoints/clan?id=multiple-clans
func (c *Client) Clans(tags []string, params url.Values) (clans []Clan, err error) {
	var b []byte
	path := "/clan/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clans)
	}
	return
}

// ClanBattles returns battles played by people in the specified clan.
// https://docs.royaleapi.com/#/endpoints/clan_battles
func (c *Client) ClanBattles(tag string, params url.Values) (battles []Battle, err error) {
	var b []byte
	path := "/clan/" + tag + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// ClanWar returns data about the current clan war.
// https://docs.royaleapi.com/#/endpoints/clan_war
func (c *Client) ClanWar(tag string, params url.Values) (war ClanWar, err error) {
	var b []byte
	path := "/clan/" + tag + "/war"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &war)
	}
	return
}

// ClanWarLog returns data about past clan wars.
// https://docs.royaleapi.com/#/endpoints/clan_warlog
func (c *Client) ClanWarLog(tag string, params url.Values) (warlog []ClanWarLogEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/warlog"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &warlog)
	}
	return
}

// ClanHistory returns a time series of member stats.
// This will only work with clans that have enabled stat tracking.
// https://docs.royaleapi.com/#/endpoints/clan_history
func (c *Client) ClanHistory(tag string, params url.Values) (history []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &history)
	}
	return
}

// ClanWeeklyHistory works like ClanHistory but returns weekly stats.
// https://docs.royaleapi.com/#/endpoints/clan_history_weekly
func (c *Client) ClanWeeklyHistory(tag string, params url.Values) (history []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history/weekly"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &history)
	}
	return
}

// ClanTracking returns basic data on whether a clan is tracked.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
func (c *Client) ClanTracking(tag string, params url.Values) (tracking ClanTracking, err error) {
	var b []byte
	path := "/clan/" + tag + "/tracking"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tracking)
	}
	return
}

// OpenTournaments returns a slice of open tournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments_open
func (c *Client) OpenTournaments(params url.Values) (tournaments []Tournament, err error) {
	var b []byte
	path := "/tournaments/open"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// KnownTournaments returns a slice of tournaments people have searched for.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
func (c *Client) KnownTournaments(params url.Values) (tournaments []Tournament, err error) {
	var b []byte
	path := "/tournaments/known"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// Get1kTournaments returns a slice of tournaments that have 1000 MaxPlayers.
func (c *Client) Get1kTournaments(params url.Values) (tournaments []Tournament1k, err error) {
	var b []byte
	path := "/tournaments/1k"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// PrepTournaments returns a slice of tournaments that have a Status of "inPreparation".
func (c *Client) PrepTournaments(params url.Values) (tournaments []PrepTournament, err error) {
	var b []byte
	path := "/tournaments/prep"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// TournamentSearch returns a slice of tournaments by a name to search for.
// https://docs.royaleapi.com/#/endpoints/tournaments_search
func (c *Client) TournamentSearch(params url.Values) (tournaments []SearchedTournament, err error) {
	var b []byte
	path := "/tournaments/search"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// Tournament returns the specified Tournament by tag.
// https://docs.royaleapi.com/#/endpoints/tournaments
func (c *Client) Tournament(tag string, params url.Values) (tournament SpecificTournament, err error) {
	var b []byte
	path := "/tournaments/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournament)
	}
	return
}

// Tournaments works like Tournament but can return multiple SpecificTournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments?id=multiple-tournaments
func (c *Client) Tournaments(tags []string, params url.Values) (tournaments []SpecificTournament, err error) {
	var b []byte
	path := "/tournaments/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// TopClans returns the top 200 clans of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_clans
func (c *Client) TopClans(location string, params url.Values) (topClans []TopClan, err error) {
	var b []byte
	path := "/top/clans/" + location
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &topClans)
	}
	return
}

// TopPlayers returns the top 200 players of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_players
func (c *Client) TopPlayers(location string, params url.Values) (topPlayers []TopPlayer, err error) {
	var b []byte
	path := "/top/players/" + location
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &topPlayers)
	}
	return
}

// PopularClans returns stats on how often a clan's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_clans
func (c *Client) PopularClans(params url.Values) (popularClans []PopularClan, err error) {
	var b []byte
	path := "/popular/clans"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularClans)
	}
	return
}

// PopularPlayers returns stats on how often a player's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_players
func (c *Client) PopularPlayers(params url.Values) (popularPlayers []PopularPlayer, err error) {
	var b []byte
	path := "/popular/players"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularPlayers)
	}
	return
}

// PopularTournaments returns stats on how often a tournament's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_tournaments
func (c *Client) PopularTournaments(params url.Values) (popularTournaments []PopularTournament, err error) {
	var b []byte
	path := "/popular/tournaments"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularTournaments)
	}
	return
}

// PopularDecks returns stats on how often a deck's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_decks
func (c *Client) PopularDecks(params url.Values) (popularDecks []PopularDeck, err error) {
	var b []byte
	path := "/popular/decks"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularDecks)
	}
	return
}

// APIKeyStats returns information about the currently authenticated token.
// https://docs.royaleapi.com/#/endpoints/auth_stats
func (c *Client) APIKeyStats(params url.Values) (keyStats APIKeyStats, err error) {
	var b []byte
	path := "/auth/stats"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &keyStats)
	}
	return
}

// Endpoints returns all the available endpoints for the API.
// It does not have any special incorporation with this wrapper and is simply included for completion's sake.
// https://docs.royaleapi.com/#/endpoints/endpoints
func (c *Client) Endpoints(params url.Values) (endpoints []string, err error) {
	var b []byte
	path := "/endpoints"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &endpoints)
	}
	return
}
