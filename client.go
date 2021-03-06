package tatsumakigo

import "context"

// Client is the main interface to interact with the API.
type Client struct {
	restClient *restClient
}

// New creates a new instance of the Tatsumaki client.
func New(token string) *Client {
	return &Client{
		newRestClient(token),
	}
}

// AdjustGuildUserPoints wraps AdjustGuildUserPointsWithContext using the background context.
func (t *Client) AdjustGuildUserPoints(guildID string, userID string, amount int, action Action) (*GuildUserPoints, error) {
	return t.AdjustGuildUserPointsWithContext(context.Background(), guildID, userID, amount, action)
}

// AdjustGuildUserPointsWithContext adjusts a user's points in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserPointsWithContext(ctx context.Context, guildID string, userID string, amount int,
	action Action) (*GuildUserPoints, error) {
	return t.restClient.adjustGuildUserPoints(ctx, guildID, userID, amount, action)
}

// AdjustGuildUserScore wraps AdjustGuildUserScoreWithContext using the background context.
func (t *Client) AdjustGuildUserScore(guildID string, userID string, amount int,
	action Action) (*GuildUserScore, error) {
	return t.AdjustGuildUserScoreWithContext(context.Background(), guildID, userID, amount, action)
}

// AdjustGuildUserScoreWithContext adjusts a user's score in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserScoreWithContext(ctx context.Context, guildID string, userID string, amount int,
	action Action) (*GuildUserScore, error) {
	return t.restClient.adjustGuildUserScore(ctx, guildID, userID, amount, action)
}

// GuildLeaderboard wraps GuildLeaderboardWithContext using the background context.
// To get all guild member rankings for the leaderboard, set limit to -1.
// If limit is set to 0, the default value will be used as per the API.
func (t *Client) GuildLeaderboard(guildID string, limit int) ([]*GuildRankedUser, error) {
	return t.GuildLeaderboardWithContext(context.Background(), guildID, limit)
}

// GuildLeaderboardWithContext gets the leaderboard for a guild.
// To get all guild member rankings for the leaderboard, set limit to -1.
// If limit is set to 0, the default value will be used as per the API.
func (t *Client) GuildLeaderboardWithContext(ctx context.Context, guildID string, limit int) ([]*GuildRankedUser, error) {
	return t.restClient.guildLeaderboard(ctx, guildID, limit)
}

// GuildUserStats wraps GuildUserStatsWithContext using the background context.
func (t *Client) GuildUserStats(guildID string, userID string) (*GuildUserStats, error) {
	return t.GuildUserStatsWithContext(context.Background(), guildID, userID)
}

// GuildUserStatsWithContext gets a user's stats for a guild.
// @me is accepted for the user ID, which will retrieve stats for the owner of the API token.
func (t *Client) GuildUserStatsWithContext(ctx context.Context, guildID string, userID string) (*GuildUserStats, error) {
	return t.restClient.guildUserStats(ctx, guildID, userID)
}

// User wraps UserWithContext using the background context.
func (t *Client) User(userID string) (*User, error) {
	return t.UserWithContext(context.Background(), userID)
}

// UserWithContext gets a Tatsumaki user profile.
func (t *Client) UserWithContext(ctx context.Context, userID string) (*User, error) {
	return t.restClient.user(ctx, userID)
}
