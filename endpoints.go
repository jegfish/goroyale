package goroyale

import (
	"encoding/json"
	"net/url"
	"strings"
)

// GetAPIVersion requests the current version of the API.
// https://docs.royaleapi.com/#/endpoints/version
func (c *Client) GetAPIVersion() (ver string, err error) {
	bytes, err := c.get("/version", url.Values{})
	if err != nil {
		return
	}
	ver = string(bytes)
	return
}

// GetConstants returns constants from the API.
// https://docs.royaleapi.com/#/endpoints/constants
func (c *Client) GetConstants(params url.Values) (constants Constants, err error) {
	var b []byte
	path := "/constants"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &constants)
	}
	return
}

// GetPlayer retrieves a player by their tag.
// https://docs.royaleapi.com/#/endpoints/player
func (c *Client) GetPlayer(tag string, params url.Values) (player Player, err error) {
	var b []byte
	path := "/player/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &player)
	}
	return
}

// GetPlayers works like GetPlayer but can return multiple players.
// The API asks that you don't include more than 7 tags in this request.
// https://docs.royaleapi.com/#/endpoints/player?id=multiple-players
func (c *Client) GetPlayers(tags []string, params url.Values) (players []Player, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &players)
	}
	return
}

// GetPlayerBattles gets battles a player participated in.
// https://docs.royaleapi.com/#/endpoints/player_battles
func (c *Client) GetPlayerBattles(tag string, params url.Values) (battles []Battle, err error) {
	var b []byte
	path := "/player/" + tag + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// GetPlayersBattles works like GetPlayerBattles but can return battles from multiple players.
// https://docs.royaleapi.com/#/endpoints/player_battles?id=multiple-tags
func (c *Client) GetPlayersBattles(tags []string, params url.Values) (battles [][]Battle, err error) {
	var b []byte
	path := "/player/" + strings.Join(tags, ",") + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// GetPlayerChests gets a player's upcoming chests.
// https://docs.royaleapi.com/#/endpoints/player_chests
func (c *Client) GetPlayerChests(tag string, params url.Values) (chests PlayerChests, err error) {
	var b []byte
	path := "/player/" + tag + "/chests"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &chests)
	}
	return
}

// GetPlayersChests works like GetPlayerChests but can return chests for multiple players.
// https://docs.royaleapi.com/#/endpoints/player_chests?id=multiple-players
func (c *Client) GetPlayersChests(tags []string, params url.Values) (chests []PlayerChests, err error) {
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

// GetClan returns info about a specific clan.
// https://docs.royaleapi.com/#/endpoints/clan
func (c *Client) GetClan(tag string, params url.Values) (clan Clan, err error) {
	var b []byte
	path := "/clan/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clan)
	}
	return
}

// GetClans works like GetClan but can return multiple clans.
// https://docs.royaleapi.com/#/endpoints/clan?id=multiple-clans
func (c *Client) GetClans(tags []string, params url.Values) (clans []Clan, err error) {
	var b []byte
	path := "/clan/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clans)
	}
	return
}

// GetClanBattles returns battles played by people in the specified clan.
// https://docs.royaleapi.com/#/endpoints/clan_battles
func (c *Client) GetClanBattles(tag string, params url.Values) (battles []Battle, err error) {
	var b []byte
	path := "/clan/" + tag + "/battles"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// GetClanWar returns data about the current clan war.
// https://docs.royaleapi.com/#/endpoints/clan_war
func (c *Client) GetClanWar(tag string, params url.Values) (war ClanWar, err error) {
	var b []byte
	path := "/clan/" + tag + "/war"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &war)
	}
	return
}

// GetClanWarLog returns data about past clan wars.
// https://docs.royaleapi.com/#/endpoints/clan_warlog
func (c *Client) GetClanWarLog(tag string, params url.Values) (warlog []ClanWarLogEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/warlog"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &warlog)
	}
	return
}

// GetClanHistory returns a time series of member stats.
// This will only work with clans that have enabled stat tracking.
// https://docs.royaleapi.com/#/endpoints/clan_history
func (c *Client) GetClanHistory(tag string, params url.Values) (history []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &history)
	}
	return
}

// GetClanWeeklyHistory works like GetClanHistory but returns weekly stats.
// https://docs.royaleapi.com/#/endpoints/clan_history_weekly
func (c *Client) GetClanWeeklyHistory(tag string, params url.Values) (history []ClanHistoryEntry, err error) {
	var b []byte
	path := "/clan/" + tag + "/history/weekly"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &history)
	}
	return
}

// GetClanTracking returns basic data on whether a clan is tracked.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
func (c *Client) GetClanTracking(tag string, params url.Values) (tracking ClanTracking, err error) {
	var b []byte
	path := "/clan/" + tag + "/tracking"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tracking)
	}
	return
}

// GetOpenTournaments returns a slice of open tournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments_open
func (c *Client) GetOpenTournaments(params url.Values) (tournaments []Tournament, err error) {
	var b []byte
	path := "/tournaments/open"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// GetKnownTournaments returns a slice of tournaments people have searched for.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
func (c *Client) GetKnownTournaments(params url.Values) (tournaments []Tournament, err error) {
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

// GetPrepTournaments returns a slice of tournaments that have a Status of "inPreparation".
func (c *Client) GetPrepTournaments(params url.Values) (tournaments []PrepTournament, err error) {
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

// GetTournament returns the specified Tournament by tag.
// https://docs.royaleapi.com/#/endpoints/tournaments
func (c *Client) GetTournament(tag string, params url.Values) (tournament SpecificTournament, err error) {
	var b []byte
	path := "/tournaments/" + tag
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournament)
	}
	return
}

// GetTournaments works like GetTournament but can return multiple SpecificTournaments.
// https://docs.royaleapi.com/#/endpoints/tournaments?id=multiple-tournaments
func (c *Client) GetTournaments(tags []string, params url.Values) (tournaments []SpecificTournament, err error) {
	var b []byte
	path := "/tournaments/" + strings.Join(tags, ",")
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &tournaments)
	}
	return
}

// GetTopClans returns the top 200 clans of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_clans
func (c *Client) GetTopClans(location string, params url.Values) (topClans []TopClan, err error) {
	var b []byte
	path := "/top/clans/" + location
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &topClans)
	}
	return
}

// GetTopPlayers returns the top 200 players of a location/global leaderboard.
// https://docs.royaleapi.com/#/endpoints/top_players
func (c *Client) GetTopPlayers(location string, params url.Values) (topPlayers []TopPlayer, err error) {
	var b []byte
	path := "/top/players/" + location
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &topPlayers)
	}
	return
}

// GetPopularClans returns stats on how often a clan's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_clans
func (c *Client) GetPopularClans(params url.Values) (popularClans []PopularClan, err error) {
	var b []byte
	path := "/popular/clans"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularClans)
	}
	return
}

// GetPopularPlayers returns stats on how often a player's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_players
func (c *Client) GetPopularPlayers(params url.Values) (popularPlayers []PopularPlayer, err error) {
	var b []byte
	path := "/popular/players"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularPlayers)
	}
	return
}

// GetPopularTournaments returns stats on how often a tournament's data has been request from the API.
// https://docs.royaleapi.com/#/endpoints/popular_tournaments
func (c *Client) GetPopularTournaments(params url.Values) (popularTournaments []PopularTournament, err error) {
	var b []byte
	path := "/popular/tournaments"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularTournaments)
	}
	return
}

// GetPopularDecks returns stats on how often a deck's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_decks
func (c *Client) GetPopularDecks(params url.Values) (popularDecks []PopularDeck, err error) {
	var b []byte
	path := "/popular/decks"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &popularDecks)
	}
	return
}

// GetAPIKeyStats returns information about the currently authenticated token.
// https://docs.royaleapi.com/#/endpoints/auth_stats
func (c *Client) GetAPIKeyStats(params url.Values) (keyStats APIKeyStats, err error) {
	var b []byte
	path := "/auth/stats"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &keyStats)
	}
	return
}

// GetEndpoints returns all the available endpoints for the API.
// It does not have any special incorporation with this wrapper and is simply included for completion's sake.
// https://docs.royaleapi.com/#/endpoints/endpoints
func (c *Client) GetEndpoints(params url.Values) (endpoints []string, err error) {
	var b []byte
	path := "/endpoints"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &endpoints)
	}
	return
}
