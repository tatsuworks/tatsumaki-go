package tatsumakigo

// GuildRankedUser is the struct for a ranked user in a guild leaderboard.
type GuildRankedUser struct {
	Rank   int64  `json:"rank"`
	Score  int64  `json:"score,string"`
	UserId string `json:"user_id"`
}

// GuildUserPoints is the struct for a user's adjusted points in a guild.
type GuildUserPoints struct {
	Points int64 `json:"points,string"`
}

// GuildUserScore is the struct for a user's adjusted score in a guild.
type GuildUserScore struct {
	Score int64 `json:"score,string"`
}

// GuildUserStats is the struct for a user's stats in a guild.
type GuildUserStats struct {
	GuildId string `json:"guild_id"`
	Points  int64  `json:"points,string"`
	Score   int64  `json:"score,string"`
	UserId  int64  `json:"user_id,string"`
}

// Ping is the struct for a ping response.
type Ping struct {
	Pong bool `json:"pong"`
}

// User is the struct for a Tatsumaki user profile.
type User struct {
	AvatarURL     string
	Background    *Background
	BadgeSlots    []*BadgeSlot
	Credits       int64
	InfoBox       string
	LevelProgress *LevelProgress
	Name          string
	Rank          int64
	Reputation    int64
	Title         string
	TotalXp       int64
}

// Background is the struct for a profile background.
type Background struct {
	ImageURL string
	Name     string
}

// Badge is the struct for a profile badge.
type Badge struct {
	ImageURL string
	Name     string
}

// BadgeSlot is the struct for a profile badge slot.
// If the badge slot does not contain an equipped badge, the badge is nil.
type BadgeSlot struct {
	Badge  *Badge
	SlotNo int
}

// LevelProgress is the struct for a user's current level progress.
type LevelProgress struct {
	CurrentXp  int64
	RequiredXp int64
}

// AdjustGuildUserPoints is the struct for the request body to adjust a user's points in a guild.
type adjustGuildUserPoints struct {
	amount int    `json:"amount"`
	action string `json:"action"`
}

// adjustGuildUserScore is the struct for the request body to adjust a user's score in a guild.
type adjustGuildUserScore struct {
	amount int    `json:"amount"`
	score  string `json:"score"`
}

// TatsumakiError is the struct for a Tatsumaki API error JSON
type tatsumakiError struct {
	message string `json:"message"`
}
