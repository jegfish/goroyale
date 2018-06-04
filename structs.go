package goroyale

import (
	"encoding/json"
)

// Structs for unmarshalling JSON from API endpoints

// Player represents a player's profile with basic stats and card collection.
// https://docs.royaleapi.com/#/endpoints/player
type Player struct {
	Tag              string
	Name             string
	Trophies         int
	Rank             int // Player's global ranking
	Arena            Arena
	Clan             PlayerClan
	Stats            PlayerStats
	Games            PlayerGames
	LeagueStatistics LeagueStatistics
	DeckLink         string // Link to copy the player's deck
	CurrentDeck      []Card
	Achievements     []Achievement
}

// Arena represents a trophy range.
type Arena struct {
	Name        string
	Arena       string // Arena's level within a league ex: "League 3"
	ArenaID     int    // Arena's number in hierarchy of arenas
	TrophyLimit int    // Upper boundary of arena trophy range
}

// PlayerClan represents a player's stats within a clan.
type PlayerClan struct {
	Tag               string
	Name              string
	Role              string
	Donations         int
	DonationsReceived int
	DonationsDelta    int
	Badge             Badge
}

// Badge represents a clan's badge/picture.
type Badge struct {
	Name     string
	Category string
	ID       int
	Image    string // Link to badge image
}

type PlayerStats struct {
	TournamentCardsWon int
	MaxTrophies        int
	ThreeCrownWins     int
	CardsFound         int
	FavoriteCard       FavoriteCard
	TotalDonations     int
	ChallengeMaxWins   int
	ChallengeCardsWon  int
	Level              int
}

// FavoriteCard is part of PlayerStats.
type FavoriteCard struct {
	Name        string
	ID          int
	MaxLevel    int
	Icon        string
	Key         string
	Elixir      int
	Type        string
	Rarity      string
	Arena       int
	Description string
}

// PlayerGames is general stats on the amount and types of games a Player has played.
type PlayerGames struct {
	Total           int
	TournamentGames int
	Wins            int
	WinsPercent     float64
	Losses          int
	LossesPercent   float64
	Draws           int
	DrawsPercent    float64
}

// LeagueStatistics represents a player's season stats.
type LeagueStatistics struct {
	CurrentSeason  CurrentSeason
	PreviousSeason PreviousSeason
	BestSeason     BestSeason
}

type CurrentSeason struct {
	Rank         int
	Trophies     int
	BestTrophies int
}

type PreviousSeason struct {
	ID           string
	Trophies     int
	BestTrophies int
}

type BestSeason struct {
	ID       string
	Rank     int
	Trophies int
}

// Card represents a card from the game.
// RequiredForUpgrade will be -1 if the card is max level.
type Card struct {
	Name               string
	Level              int
	MaxLevel           int
	Count              int
	Rarity             string
	RequiredForUpgrade requiredForUpgrade
	Icon               string
	Key                string
	Elixir             int
	Type               string
	Arena              int
	Description        string
	ID                 int
}

// In the JSON requiredForUpgrade will either be an int or the string "Maxed"
// I need to have custom JSON parsing to keep it always an int
type requiredForUpgrade int

func (r *requiredForUpgrade) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		*r = -1
		return nil
	} else {
		return json.Unmarshal(b, (*int)(r))
	}
}

// Achievement represents a player's stats and progress on an achievement.
type Achievement struct {
	Name   string
	Stars  int
	Value  int
	Target int // Value you need to reach to complete the achievement
	Info   string
}
