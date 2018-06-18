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

// TournamentMember represents a member who participated in a tournament.
type TournamentMember struct {
	Tag   string
	Name  string
	Score int
}

// SearchedTournament represents a tournament that was returned from Client.TournamentSearch().
// https://docs.royaleapi.com/#/endpoints/tournaments_search
type SearchedTournament struct {
	Tag            string
	Open           bool
	Status         string
	CreatorTag     string
	Name           string
	MaxPlayers     int
	PrepTime       int
	Duration       int
	CreateTime     int
	StartTime      int
	EndTime        int
	CurrentPlayers int
	Members        []TournamentMember
}

// Tournament is a basic tournament, other Tournament structs will have more info.
// https://docs.royaleapi.com/#/endpoints/tournaments_open
type Tournament struct {
	Tag            string
	Open           bool
	Status         string
	Name           string
	Capacity       int
	CurrentPlayers int
	MaxPlayers     int
	PrepTime       int
	Duration       int
	CreateTime     int
	StartTime      int
	EndTime        int
}

// Tournament1k is a tournament returned from Get1kTournaments.
// It always has 1000 MaxPlayers.
type Tournament1k struct {
	Tournament

	UpdatedAt int
}

// PrepTournament is a tournament returned from GetPrepTournaments.
// It always has Status set to "inPreparation".
type PrepTournament struct {
	Tournament

	UpdatedAt int
}

// SpecificTournament represents a tournament retrieved by tag with extra info included.
// https://docs.royaleapi.com/#/endpoints/tournaments_known
type SpecificTournament struct {
	Tournament

	Description string
	Creator     TournamentMember
	Members     []TournamentMember
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
	Popularity     Popularity
	Tag            string
	Open           bool
	Status         string
	Name           string
	Description    string
	MaxPlayers     int
	PrepTime       int
	Duration       int
	CreateTime     int
	StartTime      int
	EndTime        int
	CurrentPlayers int
	Creator        TournamentMember
	Members        []TournamentMember
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

// Constants represents API constants.
// https://docs.royaleapi.com/#/endpoints/constants
type Constants struct {
	AllianceBadges []struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		ID       int    `json:"id"`
	} `json:"alliance_badges"`
	Arenas []struct {
		Name                       string `json:"name"`
		Arena                      int    `json:"arena"`
		ChestArena                 string `json:"chest_arena"`
		TvArena                    string `json:"tv_arena"`
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
		Key                        string `json:"key"`
		Title                      string `json:"title"`
		Subtitle                   string `json:"subtitle"`
		ArenaID                    int    `json:"arena_id"`
		// Either int or string
		LeagueID interface{} `json:"league_id"`
		ID       int         `json:"id"`
	} `json:"arenas"`
	Cards []struct {
		Key         string `json:"key"`
		Name        string `json:"name"`
		Elixir      int    `json:"elixir"`
		Type        string `json:"type"`
		Rarity      string `json:"rarity"`
		Arena       int    `json:"arena"`
		Description string `json:"description"`
		ID          int    `json:"id"`
	} `json:"cards"`
	CardsStats struct {
		Troop []struct {
			Name                         string  `json:"name"`
			Rarity                       string  `json:"rarity"`
			SightRange                   int     `json:"sight_range"`
			DeployTime                   int     `json:"deploy_time"`
			Speed                        int     `json:"speed,omitempty"`
			Hitpoints                    int     `json:"hitpoints"`
			HitSpeed                     int     `json:"hit_speed"`
			LoadTime                     int     `json:"load_time"`
			Damage                       int     `json:"damage,omitempty"`
			LoadFirstHit                 bool    `json:"load_first_hit"`
			LoadAfterRetarget            bool    `json:"load_after_retarget"`
			AllTargetsHit                bool    `json:"all_targets_hit"`
			Range                        int     `json:"range"`
			AttacksGround                bool    `json:"attacks_ground"`
			AttacksAir                   bool    `json:"attacks_air"`
			TargetOnlyBuildings          bool    `json:"target_only_buildings"`
			CrowdEffects                 bool    `json:"crowd_effects"`
			IgnorePushback               bool    `json:"ignore_pushback"`
			Scale                        int     `json:"scale"`
			CollisionRadius              int     `json:"collision_radius"`
			Mass                         int     `json:"mass"`
			ShowHealthNumber             bool    `json:"show_health_number"`
			FlyDirectPaths               bool    `json:"fly_direct_paths"`
			FlyFromGround                bool    `json:"fly_from_ground"`
			HealOnMorph                  bool    `json:"heal_on_morph"`
			MorphKeepTarget              bool    `json:"morph_keep_target"`
			DestroyAtLimit               bool    `json:"destroy_at_limit"`
			DeathSpawnPushback           bool    `json:"death_spawn_pushback"`
			DeathInheritIgnoreList       bool    `json:"death_inherit_ignore_list"`
			Kamikaze                     bool    `json:"kamikaze"`
			ProjectileStartRadius        int     `json:"projectile_start_radius,omitempty"`
			ProjectileStartZ             int     `json:"projectile_start_z,omitempty"`
			DontStopMoveAnim             bool    `json:"dont_stop_move_anim"`
			IsSummonerTower              bool    `json:"is_summoner_tower"`
			SelfAsAoeCenter              bool    `json:"self_as_aoe_center"`
			HidesWhenNotAttacking        bool    `json:"hides_when_not_attacking"`
			HideBeforeFirstHit           bool    `json:"hide_before_first_hit"`
			SpecialAttackWhenHidden      bool    `json:"special_attack_when_hidden"`
			HasRotationOnTimeline        bool    `json:"has_rotation_on_timeline"`
			JumpEnabled                  bool    `json:"jump_enabled"`
			RetargetAfterAttack          bool    `json:"retarget_after_attack"`
			BurstKeepTarget              bool    `json:"burst_keep_target"`
			BurstAffectAnimation         bool    `json:"burst_affect_animation"`
			BuildingTarget               bool    `json:"building_target"`
			SpawnConstPriority           bool    `json:"spawn_const_priority"`
			NameEn                       string  `json:"name_en"`
			Key                          string  `json:"key,omitempty"`
			Elixir                       int     `json:"elixir,omitempty"`
			Type                         string  `json:"type,omitempty"`
			Arena                        int     `json:"arena,omitempty"`
			Description                  string  `json:"description,omitempty"`
			ID                           int     `json:"id,omitempty"`
			SpeedEn                      string  `json:"speed_en"`
			Dps                          float64 `json:"dps"`
			Projectile                   string  `json:"projectile,omitempty"`
			DeployDelay                  int     `json:"deploy_delay,omitempty"`
			StopMovementAfterMs          int     `json:"stop_movement_after_ms,omitempty"`
			WaitMs                       int     `json:"wait_ms,omitempty"`
			SightClip                    int     `json:"sight_clip,omitempty"`
			SightClipSide                int     `json:"sight_clip_side,omitempty"`
			WalkingSpeedTweakPercentage  int     `json:"walking_speed_tweak_percentage,omitempty"`
			FlyingHeight                 int     `json:"flying_height,omitempty"`
			DeathSpawnCharacter          string  `json:"death_spawn_character,omitempty"`
			SpawnStartTime               int     `json:"spawn_start_time,omitempty"`
			SpawnInterval                int     `json:"spawn_interval,omitempty"`
			SpawnNumber                  int     `json:"spawn_number,omitempty"`
			SpawnPauseTime               int     `json:"spawn_pause_time,omitempty"`
			SpawnCharacterLevelIndex     int     `json:"spawn_character_level_index,omitempty"`
			SpawnCharacter               string  `json:"spawn_character,omitempty"`
			DeathDamageRadius            int     `json:"death_damage_radius,omitempty"`
			DeathDamage                  int     `json:"death_damage,omitempty"`
			DeathPushBack                int     `json:"death_push_back,omitempty"`
			DeathSpawnCount              int     `json:"death_spawn_count,omitempty"`
			DeathSpawnRadius             int     `json:"death_spawn_radius,omitempty"`
			AreaDamageRadius             int     `json:"area_damage_radius,omitempty"`
			SpawnRadius                  int     `json:"spawn_radius,omitempty"`
			ChargeRange                  int     `json:"charge_range,omitempty"`
			DamageSpecial                int     `json:"damage_special,omitempty"`
			DamageEffectSpecial          string  `json:"damage_effect_special,omitempty"`
			ChargeSpeedMultiplier        int     `json:"charge_speed_multiplier,omitempty"`
			JumpHeight                   int     `json:"jump_height,omitempty"`
			JumpSpeed                    int     `json:"jump_speed,omitempty"`
			CustomFirstProjectile        string  `json:"custom_first_projectile,omitempty"`
			MultipleProjectiles          int     `json:"multiple_projectiles,omitempty"`
			ShieldHitpoints              int     `json:"shield_hitpoints,omitempty"`
			CrownTowerDamagePercent      int     `json:"crown_tower_damage_percent,omitempty"`
			SpawnPathfindSpeed           int     `json:"spawn_pathfind_speed,omitempty"`
			AttackPushBack               int     `json:"attack_push_back,omitempty"`
			ProjectileEffectSpecial      string  `json:"projectile_effect_special,omitempty"`
			LoadAttackEffect1            string  `json:"load_attack_effect1,omitempty"`
			LoadAttackEffect2            string  `json:"load_attack_effect2,omitempty"`
			LoadAttackEffect3            string  `json:"load_attack_effect3,omitempty"`
			LoadAttackEffectReady        string  `json:"load_attack_effect_ready,omitempty"`
			RotateAngleSpeed             int     `json:"rotate_angle_speed,omitempty"`
			VariableDamage2              int     `json:"variable_damage2,omitempty"`
			VariableDamageTime1          int     `json:"variable_damage_time1,omitempty"`
			VariableDamage3              int     `json:"variable_damage3,omitempty"`
			VariableDamageTime2          int     `json:"variable_damage_time2,omitempty"`
			TargettedDamageEffect1       string  `json:"targetted_damage_effect1,omitempty"`
			TargettedDamageEffect2       string  `json:"targetted_damage_effect2,omitempty"`
			TargettedDamageEffect3       string  `json:"targetted_damage_effect3,omitempty"`
			FlameEffect1                 string  `json:"flame_effect1,omitempty"`
			FlameEffect2                 string  `json:"flame_effect2,omitempty"`
			FlameEffect3                 string  `json:"flame_effect3,omitempty"`
			TargetEffectY                int     `json:"target_effect_y,omitempty"`
			VisualHitSpeed               int     `json:"visual_hit_speed,omitempty"`
			SpawnDeployBaseAnim          string  `json:"spawn_deploy_base_anim,omitempty"`
			SpawnAngleShift              int     `json:"spawn_angle_shift,omitempty"`
			DeathSpawnDeployTime         int     `json:"death_spawn_deploy_time,omitempty"`
			AttackShakeTime              int     `json:"attack_shake_time,omitempty"`
			MultipleTargets              int     `json:"multiple_targets,omitempty"`
			BuffOnDamage                 string  `json:"buff_on_damage,omitempty"`
			BuffOnDamageTime             int     `json:"buff_on_damage_time,omitempty"`
			SpawnAreaObject              string  `json:"spawn_area_object,omitempty"`
			SpawnAreaObjectLevelIndex    int     `json:"spawn_area_object_level_index,omitempty"`
			DashImmuneToDamageTime       int     `json:"dash_immune_to_damage_time,omitempty"`
			DashCooldown                 int     `json:"dash_cooldown,omitempty"`
			DashDamage                   int     `json:"dash_damage,omitempty"`
			DashFilter                   string  `json:"dash_filter,omitempty"`
			DashMinRange                 int     `json:"dash_min_range,omitempty"`
			DashMaxRange                 int     `json:"dash_max_range,omitempty"`
			HideTimeMs                   int     `json:"hide_time_ms,omitempty"`
			BuffWhenNotAttacking         string  `json:"buff_when_not_attacking,omitempty"`
			BuffWhenNotAttackingTime     int     `json:"buff_when_not_attacking_time,omitempty"`
			AttachedCharacter            string  `json:"attached_character,omitempty"`
			TargetedEffectVisualPushback int     `json:"targeted_effect_visual_pushback,omitempty"`
			AttackDashTime               int     `json:"attack_dash_time,omitempty"`
			LoopingFilter                string  `json:"looping_filter,omitempty"`
			LifeTime                     int     `json:"life_time,omitempty"`
			MorphTime                    int     `json:"morph_time,omitempty"`
			DashPushBack                 int     `json:"dash_push_back,omitempty"`
			DashRadius                   int     `json:"dash_radius,omitempty"`
			DashConstantTime             int     `json:"dash_constant_time,omitempty"`
			DashLandingTime              int     `json:"dash_landing_time,omitempty"`
			SpawnLimit                   int     `json:"spawn_limit,omitempty"`
			SpawnPushback                int     `json:"spawn_pushback,omitempty"`
			SpawnPushbackRadius          int     `json:"spawn_pushback_radius,omitempty"`
			KamikazeTime                 int     `json:"kamikaze_time,omitempty"`
		} `json:"troop"`
		Building []struct {
			Name                          string `json:"name"`
			Rarity                        string `json:"rarity"`
			SightRange                    int    `json:"sight_range,omitempty"`
			Hitpoints                     int    `json:"hitpoints,omitempty"`
			HitSpeed                      int    `json:"hit_speed,omitempty"`
			LoadTime                      int    `json:"load_time,omitempty"`
			LoadFirstHit                  bool   `json:"load_first_hit"`
			LoadAfterRetarget             bool   `json:"load_after_retarget"`
			Projectile                    string `json:"projectile,omitempty"`
			AllTargetsHit                 bool   `json:"all_targets_hit"`
			Range                         int    `json:"range,omitempty"`
			AttacksGround                 bool   `json:"attacks_ground"`
			AttacksAir                    bool   `json:"attacks_air"`
			TargetOnlyBuildings           bool   `json:"target_only_buildings"`
			AttachedCharacterHeight       int    `json:"attached_character_height,omitempty"`
			CrowdEffects                  bool   `json:"crowd_effects"`
			IgnorePushback                bool   `json:"ignore_pushback"`
			Scale                         int    `json:"scale"`
			CollisionRadius               int    `json:"collision_radius,omitempty"`
			TileSizeOverride              int    `json:"tile_size_override,omitempty"`
			ShowHealthNumber              bool   `json:"show_health_number"`
			FlyDirectPaths                bool   `json:"fly_direct_paths"`
			FlyFromGround                 bool   `json:"fly_from_ground"`
			HealOnMorph                   bool   `json:"heal_on_morph"`
			MorphKeepTarget               bool   `json:"morph_keep_target"`
			DestroyAtLimit                bool   `json:"destroy_at_limit"`
			DeathSpawnPushback            bool   `json:"death_spawn_pushback"`
			DeathInheritIgnoreList        bool   `json:"death_inherit_ignore_list"`
			Kamikaze                      bool   `json:"kamikaze"`
			ProjectileStartRadius         int    `json:"projectile_start_radius,omitempty"`
			ProjectileStartZ              int    `json:"projectile_start_z,omitempty"`
			DontStopMoveAnim              bool   `json:"dont_stop_move_anim"`
			IsSummonerTower               bool   `json:"is_summoner_tower"`
			NoDeploySizeW                 int    `json:"no_deploy_size_w,omitempty"`
			NoDeploySizeH                 int    `json:"no_deploy_size_h,omitempty"`
			SelfAsAoeCenter               bool   `json:"self_as_aoe_center"`
			HidesWhenNotAttacking         bool   `json:"hides_when_not_attacking"`
			HideBeforeFirstHit            bool   `json:"hide_before_first_hit"`
			SpecialAttackWhenHidden       bool   `json:"special_attack_when_hidden"`
			HasRotationOnTimeline         bool   `json:"has_rotation_on_timeline"`
			TurretMovement                int    `json:"turret_movement,omitempty"`
			ProjectileYOffset             int    `json:"projectile_y_offset,omitempty"`
			JumpEnabled                   bool   `json:"jump_enabled"`
			RetargetAfterAttack           bool   `json:"retarget_after_attack"`
			BurstKeepTarget               bool   `json:"burst_keep_target"`
			BurstAffectAnimation          bool   `json:"burst_affect_animation"`
			BuildingTarget                bool   `json:"building_target"`
			SpawnConstPriority            bool   `json:"spawn_const_priority"`
			NameEn                        string `json:"name_en,omitempty"`
			AttachedCharacter             string `json:"attached_character,omitempty"`
			DeployTime                    int    `json:"deploy_time,omitempty"`
			LifeTime                      int    `json:"life_time,omitempty"`
			Key                           string `json:"key,omitempty"`
			Elixir                        int    `json:"elixir,omitempty"`
			Type                          string `json:"type,omitempty"`
			Arena                         int    `json:"arena,omitempty"`
			Description                   string `json:"description,omitempty"`
			ID                            int    `json:"id,omitempty"`
			SpawnNumber                   int    `json:"spawn_number,omitempty"`
			SpawnPauseTime                int    `json:"spawn_pause_time,omitempty"`
			SpawnCharacterLevelIndex      int    `json:"spawn_character_level_index,omitempty"`
			SpawnCharacter                string `json:"spawn_character,omitempty"`
			MinimumRange                  int    `json:"minimum_range,omitempty"`
			Damage                        int    `json:"damage,omitempty"`
			VariableDamage2               int    `json:"variable_damage2,omitempty"`
			VariableDamageTime1           int    `json:"variable_damage_time1,omitempty"`
			VariableDamage3               int    `json:"variable_damage3,omitempty"`
			VariableDamageTime2           int    `json:"variable_damage_time2,omitempty"`
			TargettedDamageEffect1        string `json:"targetted_damage_effect1,omitempty"`
			TargettedDamageEffect2        string `json:"targetted_damage_effect2,omitempty"`
			TargettedDamageEffect3        string `json:"targetted_damage_effect3,omitempty"`
			DamageLevelTransitionEffect12 string `json:"damage_level_transition_effect12,omitempty"`
			DamageLevelTransitionEffect23 string `json:"damage_level_transition_effect23,omitempty"`
			FlameEffect1                  string `json:"flame_effect1,omitempty"`
			FlameEffect2                  string `json:"flame_effect2,omitempty"`
			FlameEffect3                  string `json:"flame_effect3,omitempty"`
			TargetEffectY                 int    `json:"target_effect_y,omitempty"`
			SpawnInterval                 int    `json:"spawn_interval,omitempty"`
			HideTimeMs                    int    `json:"hide_time_ms,omitempty"`
			UpTimeMs                      int    `json:"up_time_ms,omitempty"`
			ManaCollectAmount             int    `json:"mana_collect_amount,omitempty"`
			ManaGenerateTimeMs            int    `json:"mana_generate_time_ms,omitempty"`
			DeathSpawnCount               int    `json:"death_spawn_count,omitempty"`
			DeathSpawnCharacter           string `json:"death_spawn_character,omitempty"`
			DeathDamageRadius             int    `json:"death_damage_radius,omitempty"`
			DeathDamage                   int    `json:"death_damage,omitempty"`
			DeathPushBack                 int    `json:"death_push_back,omitempty"`
			DeathSpawnRadius              int    `json:"death_spawn_radius,omitempty"`
			DeathSpawnMinRadius           int    `json:"death_spawn_min_radius,omitempty"`
			DeathSpawnDeployTime          int    `json:"death_spawn_deploy_time,omitempty"`
		} `json:"building"`
		Spell []struct {
			Name                                   string `json:"name"`
			Rarity                                 string `json:"rarity"`
			LifeDuration                           int    `json:"life_duration"`
			LifeDurationIncreasePerLevel           int    `json:"life_duration_increase_per_level"`
			LifeDurationIncreaseAfterTournamentCap int    `json:"life_duration_increase_after_tournament_cap"`
			AffectsHidden                          bool   `json:"affects_hidden"`
			Radius                                 int    `json:"radius"`
			HitSpeed                               int    `json:"hit_speed"`
			Damage                                 int    `json:"damage"`
			NoEffectToCrownTowers                  bool   `json:"no_effect_to_crown_towers"`
			CrownTowerDamagePercent                int    `json:"crown_tower_damage_percent"`
			HitBiggestTargets                      bool   `json:"hit_biggest_targets"`
			Buff                                   string `json:"buff"`
			BuffTime                               int    `json:"buff_time"`
			BuffTimeIncreasePerLevel               int    `json:"buff_time_increase_per_level"`
			BuffTimeIncreaseAfterTournamentCap     int    `json:"buff_time_increase_after_tournament_cap"`
			CapBuffTimeToAreaEffectTime            bool   `json:"cap_buff_time_to_area_effect_time"`
			BuffNumber                             int    `json:"buff_number"`
			OnlyEnemies                            bool   `json:"only_enemies"`
			OnlyOwnTroops                          bool   `json:"only_own_troops"`
			IgnoreBuildings                        bool   `json:"ignore_buildings"`
			IgnoreHero                             bool   `json:"ignore_hero"`
			Projectile                             string `json:"projectile"`
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
			Key                                    string `json:"key,omitempty"`
			Elixir                                 int    `json:"elixir,omitempty"`
			Type                                   string `json:"type,omitempty"`
			Arena                                  int    `json:"arena,omitempty"`
			Description                            string `json:"description,omitempty"`
			ID                                     int    `json:"id,omitempty"`
		} `json:"spell"`
	} `json:"cards_stats"`
	Challenges []struct {
		Name                string `json:"name"`
		GameMode            string `json:"game_mode"`
		Enabled             bool   `json:"enabled"`
		JoinCost            int    `json:"join_cost"`
		JoinCostResource    string `json:"join_cost_resource"`
		MaxWins             int    `json:"max_wins"`
		MaxLoss             int    `json:"max_loss"`
		RewardCards         []int  `json:"reward_cards"`
		RewardGold          []int  `json:"reward_gold"`
		RewardSpell         string `json:"reward_spell"`
		RewardSpellMaxCount int    `json:"reward_spell_max_count"`
		NameEn              string `json:"name_en,omitempty"`
		Key                 string `json:"key"`
		ID                  int    `json:"id"`
	} `json:"challenges"`
	ChestOrder struct {
		MainCycle       []string `json:"MainCycle"`
		QuestEarlygame1 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_earlygame_1"`
		QuestEarlygame2 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_earlygame_2"`
		QuestLategame1 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_1"`
		QuestLategame2 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_2"`
		QuestLategame3 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_3"`
		QuestLategame4 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_4"`
		QuestLategame5 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_5"`
		QuestLategame6 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_6"`
		QuestLategame7 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_7"`
		QuestLategame8 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_8"`
		QuestLategame9 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_9"`
		QuestLategame10 []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_lategame_10"`
		QuestArena3Super []struct {
			Chest          string `json:"chest"`
			ArenaThreshold string `json:"arena_threshold"`
			OneTime        bool   `json:"one_time"`
		} `json:"Quest_arena3_super"`
	} `json:"chest_order"`
	ClanChest struct {
		OneV1 struct {
			Thresholds []int `json:"thresholds"`
			Gold       []int `json:"gold"`
			Cards      []int `json:"cards"`
		} `json:"1v1"`
		TwoV2 struct {
			Thresholds []int `json:"thresholds"`
			Gold       []int `json:"gold"`
			Cards      []int `json:"cards"`
		} `json:"2v2"`
	} `json:"clan_chest"`
	GameModes []struct {
		Name                               string `json:"name"`
		CardLevelAdjustment                string `json:"card_level_adjustment"`
		DeckSelection                      string `json:"deck_selection"`
		OvertimeSeconds                    int    `json:"overtime_seconds"`
		PredefinedDecks                    string `json:"predefined_decks,omitempty"`
		SameDeckOnBoth                     bool   `json:"same_deck_on_both"`
		SeparateTeamDecks                  bool   `json:"separate_team_decks"`
		SwappingTowers                     bool   `json:"swapping_towers"`
		UseStartingElixir                  bool   `json:"use_starting_elixir"`
		Heroes                             bool   `json:"heroes"`
		Players                            string `json:"players"`
		GivesClanScore                     bool   `json:"gives_clan_score"`
		FixedDeckOrder                     bool   `json:"fixed_deck_order"`
		BattleStartCooldown                int    `json:"battle_start_cooldown,omitempty"`
		ID                                 int    `json:"id"`
		NameEn                             string `json:"name_en"`
		ElixirProductionMultiplier         int    `json:"elixir_production_multiplier,omitempty"`
		StartingElixir                     int    `json:"starting_elixir,omitempty"`
		ClanWarDescription                 string `json:"clan_war_description,omitempty"`
		ForcedDeckCards                    string `json:"forced_deck_cards,omitempty"`
		ElixirProductionOvertimeMultiplier int    `json:"elixir_production_overtime_multiplier,omitempty"`
		EventDeckSetLimit                  string `json:"event_deck_set_limit,omitempty"`
		GoldPerTower1                      int    `json:"gold_per_tower1,omitempty"`
		GoldPerTower2                      int    `json:"gold_per_tower2,omitempty"`
		GoldPerTower3                      int    `json:"gold_per_tower3,omitempty"`
		TargetTouchdowns                   int    `json:"target_touchdowns,omitempty"`
		SkinSet                            string `json:"skin_set,omitempty"`
		FixedArena                         string `json:"fixed_arena,omitempty"`
		GemsPerTower1                      int    `json:"gems_per_tower1,omitempty"`
		GemsPerTower2                      int    `json:"gems_per_tower2,omitempty"`
		GemsPerTower3                      int    `json:"gems_per_tower3,omitempty"`
	} `json:"game_modes"`
	Rarities []struct {
		Name                 string `json:"name"`
		LevelCount           int    `json:"level_count"`
		RelativeLevel        int    `json:"relative_level"`
		MirrorRelativeLevel  int    `json:"mirror_relative_level"`
		CloneRelativeLevel   int    `json:"clone_relative_level"`
		DonateCapacity       int    `json:"donate_capacity"`
		SortCapacity         int    `json:"sort_capacity"`
		DonateReward         int    `json:"donate_reward"`
		DonateXp             int    `json:"donate_xp"`
		GoldConversionValue  int    `json:"gold_conversion_value"`
		ChanceWeight         int    `json:"chance_weight"`
		BalanceMultiplier    int    `json:"balance_multiplier"`
		UpgradeExp           []int  `json:"upgrade_exp"`
		UpgradeMaterialCount []int  `json:"upgrade_material_count"`
		UpgradeCost          []int  `json:"upgrade_cost"`
		PowerLevelMultiplier []int  `json:"power_level_multiplier"`
		RefundGems           int    `json:"refund_gems"`
	} `json:"rarities"`
	Regions []struct {
		ID        int    `json:"id"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		IsCountry bool   `json:"isCountry"`
	} `json:"regions"`
	Tournaments []struct {
		CreateCost int    `json:"create_cost"`
		MaxPlayers int    `json:"max_players"`
		Key        string `json:"key"`
		Prizes     []struct {
			Rank  int `json:"rank"`
			Cards int `json:"cards"`
			Tier  int `json:"tier"`
		} `json:"prizes"`
		Cards []int `json:"cards"`
	} `json:"tournaments"`
	TreasureChests struct {
		Cycle []struct {
			Name      string      `json:"name"`
			BaseChest interface{} `json:"base_chest"`
			Arena     struct {
				Name                      string `json:"name"`
				Arena                     int    `json:"arena"`
				ChestRewardMultiplier     int    `json:"chest_reward_multiplier"`
				ShopChestRewardMultiplier int    `json:"shop_chest_reward_multiplier"`
				Key                       string `json:"key"`
				Title                     string `json:"title"`
				Subtitle                  string `json:"subtitle"`
			} `json:"arena"`
			InShop                  bool        `json:"in_shop"`
			InArenaInfo             bool        `json:"in_arena_info"`
			TournamentChest         bool        `json:"tournament_chest"`
			SurvivalChest           bool        `json:"survival_chest"`
			ShopPriceWithoutSpeedUp int         `json:"shop_price_without_speed_up"`
			TimeTakenDays           int         `json:"time_taken_days"`
			TimeTakenHours          int         `json:"time_taken_hours"`
			TimeTakenMinutes        int         `json:"time_taken_minutes"`
			TimeTakenSeconds        int         `json:"time_taken_seconds"`
			RandomSpells            int         `json:"random_spells"`
			DifferentSpells         int         `json:"different_spells"`
			ChestCountInChestCycle  int         `json:"chest_count_in_chest_cycle"`
			RareChance              int         `json:"rare_chance"`
			EpicChance              int         `json:"epic_chance"`
			LegendaryChance         int         `json:"legendary_chance"`
			SkinChance              int         `json:"skin_chance"`
			GuaranteedSpells        interface{} `json:"guaranteed_spells"`
			MinGoldPerCard          int         `json:"min_gold_per_card"`
			MaxGoldPerCard          int         `json:"max_gold_per_card"`
			SpellSet                interface{} `json:"spell_set"`
			Exp                     int         `json:"exp"`
			SortValue               int         `json:"sort_value"`
			SpecialOffer            bool        `json:"special_offer"`
			DraftChest              bool        `json:"draft_chest"`
			BoostedChest            bool        `json:"boosted_chest"`
			LegendaryOverrideChance int         `json:"legendary_override_chance"`
			Description             string      `json:"description"`
			Notification            string      `json:"notification"`
			CardCount               int         `json:"card_count"`
			MinGold                 int         `json:"min_gold"`
			MaxGold                 int         `json:"max_gold"`
			Arenas                  []struct {
				Name                      string  `json:"name"`
				Arena                     int     `json:"arena"`
				ChestRewardMultiplier     int     `json:"chest_reward_multiplier"`
				ShopChestRewardMultiplier int     `json:"shop_chest_reward_multiplier"`
				Key                       string  `json:"key"`
				Title                     string  `json:"title"`
				Subtitle                  string  `json:"subtitle"`
				CardCoundByArena          float64 `json:"card_count_by_arena"`
				CardCountCommon           float64 `json:"card_count_common"`
				CardCountRare             float64 `json:"card_count_rare"`
				CardCountEpic             float64 `json:"card_count_epic"`
				CardCountLegendary        float64 `json:"card_count_legendary"`
			} `json:"arenas"`
		} `json:"cycle"`
		Crown []struct {
			Name      string      `json:"name"`
			BaseChest interface{} `json:"base_chest"`
			Arena     struct {
				Name                      string `json:"name"`
				Arena                     int    `json:"arena"`
				ChestRewardMultiplier     int    `json:"chest_reward_multiplier"`
				ShopChestRewardMultiplier int    `json:"shop_chest_reward_multiplier"`
				Key                       string `json:"key"`
				Title                     string `json:"title"`
				Subtitle                  string `json:"subtitle"`
			} `json:"arena"`
			InShop                  bool          `json:"in_shop"`
			InArenaInfo             bool          `json:"in_arena_info"`
			TournamentChest         bool          `json:"tournament_chest"`
			SurvivalChest           bool          `json:"survival_chest"`
			ShopPriceWithoutSpeedUp int           `json:"shop_price_without_speed_up"`
			TimeTakenDays           int           `json:"time_taken_days"`
			TimeTakenHours          int           `json:"time_taken_hours"`
			TimeTakenMinutes        int           `json:"time_taken_minutes"`
			TimeTakenSeconds        int           `json:"time_taken_seconds"`
			RandomSpells            int           `json:"random_spells"`
			DifferentSpells         int           `json:"different_spells"`
			ChestCountInChestCycle  int           `json:"chest_count_in_chest_cycle"`
			RareChance              int           `json:"rare_chance"`
			EpicChance              int           `json:"epic_chance"`
			LegendaryChance         int           `json:"legendary_chance"`
			SkinChance              int           `json:"skin_chance"`
			GuaranteedSpells        interface{}   `json:"guaranteed_spells"`
			MinGoldPerCard          int           `json:"min_gold_per_card"`
			MaxGoldPerCard          int           `json:"max_gold_per_card"`
			SpellSet                interface{}   `json:"spell_set"`
			Exp                     int           `json:"exp"`
			SortValue               int           `json:"sort_value"`
			SpecialOffer            bool          `json:"special_offer"`
			DraftChest              bool          `json:"draft_chest"`
			BoostedChest            bool          `json:"boosted_chest"`
			LegendaryOverrideChance int           `json:"legendary_override_chance"`
			Description             string        `json:"description"`
			Notification            string        `json:"notification"`
			CardCount               int           `json:"card_count"`
			MinGold                 int           `json:"min_gold"`
			MaxGold                 int           `json:"max_gold"`
			Arenas                  []interface{} `json:"arenas"`
		} `json:"crown"`
		Shop []struct {
			Name      string      `json:"name"`
			BaseChest interface{} `json:"base_chest"`
			Arena     struct {
				Name                      string `json:"name"`
				Arena                     int    `json:"arena"`
				ChestRewardMultiplier     int    `json:"chest_reward_multiplier"`
				ShopChestRewardMultiplier int    `json:"shop_chest_reward_multiplier"`
				Key                       string `json:"key"`
				Title                     string `json:"title"`
				Subtitle                  string `json:"subtitle"`
			} `json:"arena"`
			InShop                  bool          `json:"in_shop"`
			InArenaInfo             bool          `json:"in_arena_info"`
			TournamentChest         bool          `json:"tournament_chest"`
			SurvivalChest           bool          `json:"survival_chest"`
			ShopPriceWithoutSpeedUp int           `json:"shop_price_without_speed_up"`
			TimeTakenDays           int           `json:"time_taken_days"`
			TimeTakenHours          int           `json:"time_taken_hours"`
			TimeTakenMinutes        int           `json:"time_taken_minutes"`
			TimeTakenSeconds        int           `json:"time_taken_seconds"`
			RandomSpells            int           `json:"random_spells"`
			DifferentSpells         int           `json:"different_spells"`
			ChestCountInChestCycle  int           `json:"chest_count_in_chest_cycle"`
			RareChance              int           `json:"rare_chance"`
			EpicChance              int           `json:"epic_chance"`
			LegendaryChance         int           `json:"legendary_chance"`
			SkinChance              int           `json:"skin_chance"`
			GuaranteedSpells        interface{}   `json:"guaranteed_spells"`
			MinGoldPerCard          int           `json:"min_gold_per_card"`
			MaxGoldPerCard          int           `json:"max_gold_per_card"`
			SpellSet                interface{}   `json:"spell_set"`
			Exp                     int           `json:"exp"`
			SortValue               int           `json:"sort_value"`
			SpecialOffer            bool          `json:"special_offer"`
			DraftChest              bool          `json:"draft_chest"`
			BoostedChest            bool          `json:"boosted_chest"`
			LegendaryOverrideChance int           `json:"legendary_override_chance"`
			Description             string        `json:"description"`
			CardCount               int           `json:"card_count"`
			MinGold                 int           `json:"min_gold"`
			MaxGold                 int           `json:"max_gold"`
			Arenas                  []interface{} `json:"arenas"`
		} `json:"shop"`
	} `json:"treasure_chests"`
}
