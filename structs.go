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
	}
	return json.Unmarshal(b, (*int)(r))
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

// PlayerChests represents info on upcoming chests for a player.
// https://docs.royaleapi.com/#/endpoints/player_chests
type PlayerChests struct {
	Upcoming     []string
	SuperMagical int
	Magical      int
	Legendary    int
	Epic         int
	Giant        int
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

// ClanWarLogEntry represents a clan war returned from the clan warlog endpoint
type ClanWarLogEntry struct {
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

// Tracking represents info on if a clan is tracked by the API or not.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
type Tracking struct {
	Active        bool
	Available     bool
	SnapshotCount int
}

// ClanTracking represents basic info on whether a clan is tracked by the API.
// https://docs.royaleapi.com/#/endpoints/clan_tracking
type ClanTracking struct {
	Tracking

	Tag string
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

// TopClan is a clan from the leaderboards.
// https://docs.royaleapi.com/#/endpoints/top_clans
type TopClan struct {
	Tag          string
	Name         string
	Score        int
	MemberCount  int
	Rank         int
	PreviousRank int
	Badge        Badge
	Location     Location
}

// TopPlayer is a player from the leaderboards.
// https://docs.royaleapi.com/#/endpoints/top_players
type TopPlayer struct {
	Name           string
	Tag            string
	Rank           int
	PreviousRank   int
	EXPLevel       int
	Trophies       int
	DonationsDelta int
	Clan           TeamClan
	Arena          Arena
}

// Popularity represents how popular an item is.
type Popularity struct {
	Hits          string
	HitsPerDayAvg float64
}

// PopularClan represents data on how often a clan has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_clans
type PopularClan struct {
	Popularity    Popularity
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
	Tracking      Tracking
}

// PopularPlayer represents data on how often a player has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_players
type PopularPlayer struct {
	Popularity   Popularity
	Tag          string
	Name         string
	Trophies     int
	Rank         int
	Arena        Arena
	Clan         PlayerClan
	Stats        PlayerStats
	Games        PlayerGames
	DeckLink     string
	CurrentDeck  []Card
	Cards        []Card
	Achievements []Achievement
}

// PopularTournament represents info on how often a tournament has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_tournaments
type PopularTournament struct {
	Popularity          Popularity
	Tag                 string
	Type                string
	Status              string
	Name                string
	Description         string
	MaxCapacity         int
	PreparationDuration int
	Duration            int
	CreateTime          int
	StartTime           int
	EndTime             int
	PlayerCount         int
	Creator             TournamentMember
	Members             []TournamentMember
}

// PopularDeckCard represents a card within a deck returned by the popular decks endpoint.
type PopularDeckCard struct {
	Arena       int
	Description string
	Elixir      int
	Icon        string
	ID          int
	Key         string
	MaxLevel    int
	Name        string
	Rarity      string
	Type        string
}

// PopularDeck represents info on how often a deck's data has been requested from the API.
// https://docs.royaleapi.com/#/endpoints/popular_decks
type PopularDeck struct {
	Popularity int
	Cards      []PopularDeckCard
	DeckLink   string
}

// APIKeyStats represents info on your API token.
// https://docs.royaleapi.com/#/endpoints/auth_stats
type APIKeyStats struct {
	ID           string
	LastRequest  int
	RequestCount map[string]int
}

// TournamentSearchEntry represents a tournament that was returned from Client.TournamentSearch().
// https://docs.royaleapi.com/#/endpoints/tournaments_search
type TournamentSearchEntry struct {
	Tag                 string
	Type                string
	Status              string
	CreatorTag          string
	Name                string
	MaxCapacity         int
	PreparationDuration int
	Duration            int
	CreateTime          int
	StartTime           int
	EndTime             int
	PlayerCount         int
	Members             []TournamentMember
}

// TODO: Add struct for /constant
type Constants struct {
	AllianceBadges []Badge `json:"alliance_badges"`
	Arenas         []ConstantsArena
	Cards          []ConstantsCard
	CardsStats     ConstantsCardsStats `json:"cards_stats"`
	Challenges     []ConstantsChallenges
	ChestOrder     ConstantsChestOrder `json:"chest_order"`
	ClanChest      ConstantsClanChest  `json:"clan_chest"`
	GameModes      []ConstantsGameMode `json:"game_modes"`
	Rarities       []ConstantsRarity
	Regions        []ConstantsRegion
	Tournaments    []ConstantsTournament
	TreasureChests ConstantsTreasureChests `json:"treasure_chests"`
}

type ConstantsArena struct {
	Name                       string
	Arena                      int
	ChestArena                 string `json:"chest_arena"`
	TVArena                    bool   `json:"tv_arena"`
	IsInUse                    bool   `json:"is_in_use"`
	TrainingCamp               bool   `json:"training_camp"`
	TrophyLimit                int    `json:"trophy_limit"`
	DemoteTrophyLimit          int    `json:"demote_trophy_limit"`
	SeasonTrophyReset          int    `json:"season_trophy_reset"`
	ChestRewardMultiplier      int    `json:"chest_reward_multiplier"`
	ShopChestRewardMultiplier  int    `json:"shop_chest_reward_multiplier"`
	RequestSize                int    `json:"request_size"`
	MaxDonationCountCommon     int    `json:"max_donation_count_common"`
	MaxDonationCountRare       int    `json:"max_donation_count_rare"`
	MaxDonationCountEpic       int    `json:"max_donation_count_epic"`
	MatchmakingMinTrophyDelta  int    `json:"matchmaking_min_trophy_delta"`
	MatchmakingMaxTrophyDelta  int    `json:"matchmaking_max_trophy_delta"`
	MatchmakingMaxSeconds      int    `json:"matchmaking_max_seconds"`
	DailyDonationCapacityLimit int    `json:"daily_donation_capacity_limit"`
	BattleRewardGold           int    `json:"battle_reward_gold"`
	SeasonRewardChest          string `json:"season_reward_chest"`
	QuestCycle                 string `json:"quest_cycle"`
	ForceQuestChestCycle       string `json:"force_quest_chest_cycle"`
	Key                        string
	Title                      string
	Subtitle                   string
	ArenaID                    int
	LeagueID                   int
	ID                         int
}

type ConstantsCard struct {
	Key         string
	Name        string
	Elixir      int
	Type        string
	Rarity      string
	Arena       int
	Description string
	ID          int
}

type ConstantsCardsStats struct {
	Troop    []ConstantsTroop
	Building []ConstantsBuilding
	Spell    []ConstantsSpell
}

type ConstantsTroop struct {
	Name                    string
	Rarity                  string
	SightRange              int `json:"sight_range"`
	DeployTime              int `json:"deploy_time"`
	Speed                   int
	Hitpoints               int
	HitSpeed                int `json:"hit_speed"`
	LoadTime                int `json:"load_time"`
	Damage                  int
	LoadFirstHit            bool `json:"load_first_hit"`
	LoadAfterRetarget       bool `json:"load_after_retarget"`
	AllTargetsHit           bool `json:"all_targets_hit"`
	Range                   int
	AttacksGround           bool `json:"attacks_ground"`
	AttacksAir              bool `json:"attacks_air"`
	TargetOnlyBuildings     bool `json:"target_only_buildings"`
	CrowdEffects            bool `json:"crowd_effects"`
	IgnorePushback          bool `json:"ignore_pushback"`
	Scale                   int
	CollisionRadius         int `json:"collision_radius"`
	Mass                    int
	ShowHealthNumber        bool `json:"show_health_number"`
	FlyDirectPaths          bool `json:"fly_direct_paths"`
	FlyFromGround           bool `json:"fly_from_ground"`
	HealOnMorph             bool `json:"heal_on_morph"`
	MorphKeepTarget         bool `json:"morph_keep_target"`
	DestroyAtLimit          bool `json:"destroy_at_limit"`
	DeathSpawnPushback      bool `json:"death_spawn_pushback"`
	DeathInheritIgnoreList  bool `json:"death_inherit_ignore_list"`
	Kamikaze                bool
	ProjectileStartRadius   int    `json:"projectile_start_radius"`
	ProjectileStartZ        int    `json:"projectile_start_z"`
	DontStopMoveAnim        bool   `json:"dont_stop_move_anim"`
	IsSummonerTower         bool   `json:"is_summoner_tower"`
	SelfAsAOECenter         bool   `json:"self_as_aoe_center"`
	HidesWhenNotAttacking   bool   `json:"hides_when_not_attacking"`
	HidesBeforeFirstHit     bool   `json:"hides_before_first_hit"`
	SpecialAttackWhenHidden bool   `json:"special_attack_when_hidden"`
	HasRotationOnTimeline   bool   `json:"has_rotation_on_timeline"`
	JumpEnabled             bool   `json:"jump_enabled"`
	RetargetAfterAttack     bool   `json:"retarget_after_attack"`
	BurstKeepTarget         bool   `json:"burst_keep_target"`
	BurstAffectAnimation    bool   `json:"burst_effect_animation"`
	BuildingTarget          bool   `json:"building_target"`
	SpawnConstPriority      bool   `json:"spawn_const_priority"`
	NameEN                  string `json:"name_en"`
	Key                     string
	Elixir                  int
	Type                    string
	Arena                   int
	Description             string
	ID                      int
	SpeedEN                 string `json:"speed_en"`
	DPS                     float64
}

type ConstantsBuilding struct {
	Name                    string
	Rarity                  string
	SightRange              int `json:"sight_range"`
	Hitpoints               int
	HitSpeed                int  `json:"hit_speed"`
	LoadTime                int  `json:"load_time"`
	LoadFirstHit            bool `json:"load_first_hit"`
	LoadAfterRetarget       bool `json:"load_after_retarget"`
	Projectile              string
	AllTargetsHit           bool `json:"all_targets_hit"`
	Range                   int
	AttacksGround           bool `json:"attacks_ground"`
	AttacksAir              bool `json:"attacks_air"`
	TargetOnlyBuildings     bool `json:"target_only_buildings"`
	AttachedCharacterHeight int  `json:"attached_character_height"`
	CrowdEffects            bool `json:"crowd_effects"`
	IgnorePushback          bool `json:"ignore_pushback"`
	Scale                   int
	CollisionRadius         int  `json:"collision_radius"`
	TileSizeOverride        int  `json:"tile_size_override"`
	ShowHealthNumber        bool `json:"show_health_number"`
	FlyDirectPaths          bool `json:"fly_direct_paths"`
	FlyFromGround           bool `json:"fly_from_ground"`
	HealOnMorph             bool `json:"heal_on_morph"`
	MorphKeepTarget         bool `json:"morph_keep_target"`
	DestroyAtLimit          bool `json:"destroy_at_limit"`
	DeathSpawnPushback      bool `json:"death_spawn_pushback"`
	DeathInheritIgnoreList  bool `json:"death_inherit_ignore_list"`
	Kamikaze                bool
	ProjectileStartRadius   int    `json:"projectile_start_radius"`
	ProjectileStartZ        int    `json:"projectile_start_z"`
	DontStopMoveAnim        bool   `json:"dont_stop_move_anim"`
	IsSummonerTower         bool   `json:"is_summoner_tower"`
	NoDeploySizeW           int    `json:"no_deploy_size_w"`
	NoDeploySizeH           int    `json:"no_deploy_size_h"`
	SelfAsAOECenter         bool   `json:"self_as_aoe_center"`
	HidesWhenNotAttacking   bool   `json:"hides_when_not_attacking"`
	HideBeforeFirstHit      bool   `json:"hide_before_first_hit"`
	SpecialAttackWhenHidden bool   `json:"special_attack_when_hidden"`
	HasRotationOnTimeline   bool   `json:"has_rotation_on_timeline"`
	TurretMovement          int    `json:"turret_movement"`
	ProjectileYOffset       int    `json:"projectile_y_offset"`
	JumpEnabled             bool   `json:"jump_enabled"`
	RetargetAfterAttack     bool   `json:"retarget_after_attack"`
	BurstKeepTarget         bool   `json:"burst_keep_target"`
	BurstAffectAnimation    bool   `json:"burst_affect_animation"`
	BuildingTarget          bool   `json:"building_target"`
	SpawnConstPriority      bool   `json:"spawn_const_priority"`
	NameEN                  string `json:"name_en"`
}

type ConstantsSpell struct {
	Name                                   string
	Rarity                                 string
	LifeDuration                           int  `json:"life_duration"`
	LifeDurationIncreasePerLevel           int  `json:"life_duration_increase_per_level"`
	LifeDurationIncreaseAfterTournamentCap int  `json:"life_duration_increase_after_tournament_cap"`
	AffectsHidden                          bool `json:"affects_hidden"`
	Radius                                 int
	HitSpeed                               int `json:"hit_speed"`
	Damage                                 int
	NoEffectToCrownTowers                  bool `json:"no_effect_to_crown_towers"`
	CrownTowerDamagePercent                int  `json:"crown_tower_damage_percent"`
	HitBiggestTargets                      bool `json:"hit_biggest_targets"`
	Buff                                   string
	BuffTime                               int  `json:"buff_time"`
	BuffTimeIncreasePerLevel               int  `json:"buff_time_increase_per_level"`
	BuffTimeIncreaseAfterTournamentCap     int  `json:"buff_time_increase_after_tournament_cap"`
	CapBuffTimeToAreaEffectTime            bool `json:"cap_buff_time_to_area_effect_time"`
	BuffNumber                             int  `json:"buff_number"`
	OnlyEnemies                            bool `json:"only_enemies"`
	OnlyOwnTroops                          bool `json:"only_own_troops"`
	IgnoreBuildings                        bool `json:"ignore_buildings"`
	IgnoreHero                             bool `json:"ignore_hero"`
	Projectile                             string
	SpawnCharacter                         string `json:"spawn_character"`
	SpawnInterval                          int    `json:"spawn_interval"`
	SpawnRandomizeSequence                 bool   `json:"spawn_randomize_sequence"`
	SpawnDeployBaseAnim                    string `json:"spawn_deploy_base_anim"`
	SpawnTime                              int    `json:"spawn_time"`
	SpawnCharacterLevelIndex               int    `json:"spawn_character_level_index"`
	SpawnInitialDelay                      int    `json:"spawn_initial_delay"`
	SpawnMaxCount                          int    `json:"spawn_max_count"`
	SpawnMaxRadius                         int    `json:"spawn_max_radius"`
	SpawnMinRadius                         int    `json:"spawn_min_radius"`
	SpawnFromMinToMax                      bool   `json:"spawn_from_min_to_max"`
	SpawnAngleShift                        int    `json:"spawn_angle_shift"`
	HitsGround                             bool   `json:"hits_ground"`
	HitsAir                                bool   `json:"hits_air"`
	Key                                    string
	Elixir                                 int
	Type                                   string
	Arena                                  int
	Description                            string
	ID                                     int
}

type ConstantsChallenges struct {
	Name                string
	GameMode            string `json:"game_mode"`
	Enabled             bool
	JoinCost            int    `json:"join_cost"`
	JoinCostResource    string `json:"join_cost_resource"`
	MaxWins             int    `json:"max_wins"`
	MaxLoss             int    `json:"max_loss"`
	RewardCards         []int  `json:"reward_cards"`
	RewardGold          []int  `json:"reward_gold"`
	RewardSpell         string `json:"reward_spell"`
	RewardSpellMaxCount int    `json:"reward_spell_max_count"`
	NameEn              string `json:"name_en"`
	Key                 string
	ID                  int
}

type ConstantsChestOrder struct {
	MainCycle        []string
	QuestEarlyGame1  []ConstantsChestQuest `json:"Quest_earlygame_1"`
	QuestEarlyGame2  []ConstantsChestQuest `json:"Quest_earlygame_2"`
	QuestLateGame1   []ConstantsChestQuest `json:"Quest_lategame_1"`
	QuestLateGame2   []ConstantsChestQuest `json:"Quest_lategame_2"`
	QuestLateGame3   []ConstantsChestQuest `json:"Quest_lategame_3"`
	QuestLateGame4   []ConstantsChestQuest `json:"Quest_lategame_4"`
	QuestLateGame5   []ConstantsChestQuest `json:"Quest_lategame_5"`
	QuestLateGame6   []ConstantsChestQuest `json:"Quest_lategame_6"`
	QuestLateGame7   []ConstantsChestQuest `json:"Quest_lategame_7"`
	QuestLateGame8   []ConstantsChestQuest `json:"Quest_lategame_8"`
	QuestLateGame9   []ConstantsChestQuest `json:"Quest_lategame_9"`
	QuestLateGame10  []ConstantsChestQuest `json:"Quest_lategame_10"`
	QuestArena3Super []ConstantsChestQuest `json:"Quest_arena3_super"`
}

type ConstantsChestQuest struct {
	Chest          string
	ArenaThreshold string `json:"arena_threshold"`
	OneTime        bool   `json:"one_time"`
}

type ConstantsClanChest struct {
	OneVOne ConstantsClanChestEntry `json:"1v1"`
	TwoVTwo ConstantsClanChestEntry `json:"2v2"`
}

type ConstantsClanChestEntry struct {
	Thresholds []int
	Gold       []int
	Cards      []int
}

type ConstantsGameMode struct {
	Name                string
	CardLevelAdjustment string `json:"card_level_adjustment"`
	DeckSelection       string `json:"deck_selection"`
	OvertimeSeconds     int    `json:"overtime_seconds"`
	PredefinedDecks     string `json:"predefined_decks"`
	SameDeckOnBoth      bool   `json:"same_deck_on_both"`
	SeparateTeamDecks   bool   `json:"separate_team_decks"`
	SwappingTowers      bool   `json:"swapping_towers"`
	UseStartingElixir   bool   `json:"use_starting_elixir"`
	Heroes              bool
	Players             string
	GivesClanScore      bool `json:"gives_clan_score"`
	FixedDeckOrder      bool `json:"fixed_deck_order"`
	BattleStartCooldown int  `json:"battle_start_cooldown"`
	ID                  int
	NameEN              string `json:"name_en"`
}

type ConstantsRarity struct {
	Name                 string
	LevelCount           int   `json:"level_count"`
	RelativeLevel        int   `json:"relative_level"`
	MirrorRelativeLevel  int   `json:"mirror_relative_level"`
	CloneRelativeLevel   int   `json:"clone_relative_level"`
	DonateCapacity       int   `json:"donate_capacity"`
	SortCapacity         int   `json:"sort_capacity"`
	DonateReward         int   `json:"donate_reward"`
	DonateXP             int   `json:"donate_xp"`
	GoldConversionValue  int   `json:"gold_conversion_value"`
	ChanceWeight         int   `json:"chance_weight"`
	BalanceMultiplier    int   `json:"balance_multiplier"`
	UpgradeEXP           []int `json:"upgrade_exp"`
	UpgradeMaterialCount []int `json:"upgrade_material_count"`
	UpgradeCost          []int `json:"upgrade_cost"`
	PowerLevelMultiplier []int `json:"power_level_multiplier"`
	RefundGems           int   `json:"refund_gems"`
}

type ConstantsRegion struct {
	ID        int
	Key       string
	Name      string
	IsCountry bool
}

type ConstantsTournament struct {
	CreateCost int `json:"create_cost"`
	MaxPlayers int `json:"max_players"`
	Key        string
	Prizes     []ConstantsTournamentPrize
	Cards      []int
}

type ConstantsTournamentPrize struct {
	Rank  int
	Cards int
	Tier  int
}

type ConstantsTreasureChests struct {
	Cycle []ConstantsTreasureChestsCycleEntry
	Crown []ConstantsTreasureChestsCrownEntry
	Shop  []ConstantsTreasureChestsShopEntry
}

type ConstantsTreasureChestsCycleEntry struct {
	Name                    string
	BaseChest               nil `json:"base_chest"`
	Arena                   ConstantsTreasureChestsArena
	InShop                  bool `json:"in_shop"`
	InArenaInfo             bool `json:"in_arena_info"`
	TournamentChest         bool `json:"tournament_chest"`
	SurvivalChest           bool `json:"survival_chest"`
	ShopPriceWithoutSpeedUp int  `json:"shop_price_without_speed_up"`
	TimeTakenDays           int  `json:"time_taken_days"`
	TimeTakenHours          int  `json:"time_taken_hours"`
	TimeTakenMinutes        int  `json:"time_taken_minutes"`
	TimeTakenSeconds        int  `json:"time_taken_seconds"`
	RandomSpells            int  `json:"random_spells"`
	DifferentSpells         int  `json:"different_spells"`
	ChestCountInChestCycle  int  `json:"chest_count_in_chest_cycle"`
	RareChance              int  `json:"rare_chance"`
	EpicChance              int  `json:"epic_chance"`
	LegendaryChance         int  `json:"legendary_chance"`
	SkinChance              int  `json:"skin_chance"`
	GuaranteedSpells        nil  `json:"guaranteed_spells"`
	MinGoldPerCard          int  `json:"min_gold_per_card"`
	MaxGoldPerCard          int  `json:"max_gold_per_card"`
	SpellSet                nil  `json:"spell_set"`
	Exp                     int
	SortValue               int  `json:"sort_value"`
	SpecialOffer            bool `json:"special_offer"`
	DraftChest              bool `json:"draft_chest"`
	BoostedChest            bool `json:"boosted_chest"`
	LegendaryOverrideChance int  `json:"legendary_override_chance"`
	Description             string
	Notification            string
	CardCount               int `json:"card_count"`
	MinGold                 int `json:"min_gold"`
	MaxGold                 int `json:"max_gold"`
	Arenas                  []ConstantsTreasureChestsArena
}

type ConstantsTreasureChestsArena struct {
	Name                      string
	Arena                     int
	ChestRewardMultiplier     int `json:"chest_reward_multiplier"`
	ShopChestRewardMultiplier int `json:"shop_chest_reward_multiplier"`
	Key                       string
	Title                     string
	Subtitle                  string
}

type ConstantsTreausureChestsCrownEntry struct {
}
