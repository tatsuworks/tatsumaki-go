package tatsumakigo

// GuildRankedUser is the struct for a ranked user in a guild leaderboard.
type GuildRankedUser struct {
	Rank   int64  `json:"rank"`
	Score  int64  `json:"score,string"`
	UserID string `json:"user_id"`
}

// GuildUserStats is the struct for a user's stats in a guild.
type GuildUserStats struct {
	GuildID string `json:"guild_id"`
	Points  int64  `json:"points,string"`
	Score   int64  `json:"score,string"`
	UserID  int64  `json:"user_id,string"`
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
	ImageUrl string
	Name     string
}

// Badge is the struct for a profile badge.
type Badge struct {
	ImageUrl string
	Name     string
}

// BadgeSlot is the struct for a profile badge slot.
// If the badge slot does not contain an equipped badge, the badge is nil.
type BadgeSlot struct {
	Badge  *Badge
	SlotNo int
}

// LevelProgress is the struct for a user's current XP progress.
type LevelProgress struct {
	CurrentXp  int64
	RequiredXp int64
}

// TatsumakiError is the struct for a Tatsumaki API error JSON
type tatsumakiError struct {
	message string `json:"message"`
}
