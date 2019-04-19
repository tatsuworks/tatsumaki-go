package tatsumaki_go

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

type Background struct {
	ImageUrl string
	Name     string
}

type Badge struct {
	ImageUrl string
	Name     string
}

type BadgeSlot struct {
	Badge  *Badge
	SlotNo int
}

type LevelProgress struct {
	CurrentXp  int64
	RequiredXp int64
}

type tatsumakiError struct {
	message string `json:"message"`
}
