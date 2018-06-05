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

// Battle represents a match played.
type Battle struct {
	Type           string
	ChallengeType  string
	Mode           BattleMode
	WinCountBefore int
	UTCTime        int
	DeckType       string
	TeamSize       int

	// Winner = TeamCrowns - OpponentCrowns
	// 0        => tie
	// positive => player won
	// negative => opponent won
	Winner int

	TeamCrowns     int
	OpponentCrowns int
	Team           []TeamMember
	Opponent       []TeamMember
	Arena          Arena
}

// BattleMode represents info on the type of battle.
type BattleMode struct {
	Name            string
	Deck            string
	CardLevels      string
	OvertimeSeconds int
	Players         string
	SameDeck        bool
}

// TeamMember represents a member of a side within a PlayerBattle
type TeamMember struct {
	Tag           string
	Name          string
	CrownsEarned  int
	TrophyChange  int
	StartTrophies int
	Clan          TeamClan
	DeckLink      string
	Deck          []Card
}

// TeamClan represents basic info on a clan within the game.
type TeamClan struct {
	Tag   string
	Name  string
	Badge Badge
}

// ClanSearch represents a clan received from the clan search endpoint.
type ClanSearch struct {
	Tag           string
	Name          string
	Type          string
	Score         int
	MemberCount   int
	RequiredScore int
	Donations     int
	Badge         Badge
	Location      Location
}

// Location represents a country.
type Location struct {
	Name      string
	IsCountry bool
	Code      string
}

// Clan represents a clan recieved directly from the clan endpoint.
type Clan struct {
	Tag           string
	Name          string
	Description   string
	Type          string
	Score         int
	MemberCount   int
	RequiredScore int
	Donations     int
	ClanChest     ClanChest
	Badge         Badge
	Location      Location
	Members       []ClanMember
}

// ClanChest is no longer in the game but the API lists it so it is here for completion's sake.
type ClanChest struct {
	Status   string
	Crowns   int
	Level    int
	MaxLevel int
}

// ClanMember represents a player inside of a clan received directly from the clan endpoint.
type ClanMember struct {
	Name              string
	Tag               string
	Rank              int // Player's ranking within the clan
	PreviousRank      int
	Role              string
	EXPLevel          int
	Trophies          int
	ClanChestCrowns   int
	Donations         int
	DonationsReceived int
	DonationsDelta    int
	DonationsPercent  float64
	Arena             Arena
}

// ClanWar represents a war a clan participated/is participating in.
type ClanWar struct {
	State             string
	WarEndTime        int
	CollectionEndTime int
	Clan              ClanWarClan
	Participants      []ClanWarParticipant
	Standings         []ClanWarClan
}

// ClanWarClan represents the clan that was queried for when getting a ClanWar.
type ClanWarClan struct {
	Tag           string
	Name          string
	Participants  int
	BattlesPlayed int
	Wins          int
	Crowns        int
	WarTrophies   int
	Badge         Badge
}

// ClanWarParticipant represents a player who was a member of a clan war.
type ClanWarParticipant struct {
	Tag           string
	Name          string
	CardsEarned   int
	BattlesPlayed int
	Wins          int
}

// ClanWarLog represents a clan war returned from the clan warlog endpoint
type ClanWarLog struct {
	CreatedDate  int
	Participants []ClanWarParticipant
	Standings    []ClanWarLogClan
	SeasonNumber int
}

// ClanWarLogClan represents a clan that participated in a clan war returned from the clan warlog endpoint.
type ClanWarLogClan struct {
	ClanWarClan
	WarTrophiesChange int
}

// ClanHistoryEntry represents a value of a key in the object returned from the clan history endpoint.
// https://docs.royaleapi.com/#/endpoints/clan_history
type ClanHistoryEntry struct {
	Donations   int
	MemberCount int
	Members     []ClanHistoryMember
}

// ClanHistoryMember represents a member of a clan within the clan history endpoint.
type ClanHistoryMember struct {
	ClanRank  int
	Crowns    int
	Donations int
	Name      string
	Tag       string
	Trophies  int
}

// ClanTracking represents basic info on whether a clan is tracked by the API.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
type ClanTracking struct {
	Tag           string
	Active        bool
	Available     bool
	SnapshotCount int
}

// OpenTournament is an open tournament.
// https://docs.royaleapi.com/#/endpoints/tournaments_open
type OpenTournament struct {
	Tag                 string
	Type                string
	Status              string
	Name                string
	Capacity            int
	PlayerCount         int
	MaxCapacity         int
	PreparationDuration int
	Duration            int
	CreateTime          int
	StartTime           int
	EndTime             int
}

// KnownTournament is a tournament someone has already searched for.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
type KnownTournament struct {
	OpenTournament
}

// Tournament represents a specific tournament with extra info included.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
type Tournament struct {
	OpenTournament

	Description string
	Creator     TournamentMember
	Members     []TournamentMember
}

// TournamentMember represents a member who participated in a tournament.
type TournamentMember struct {
	Tag   string
	Name  string
	Score int
}
