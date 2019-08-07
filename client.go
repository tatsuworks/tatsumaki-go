package tatsumakigo

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

// AdjustGuildUserPoints adjusts a user's points in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserPoints(guildID string, userID string, amount int, action Action) (*GuildUserPoints, error) {
	return t.restClient.adjustGuildUserPoints(guildID, userID, amount, action)
}

// AdjustGuildUserScore adjusts a user's score in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserScore(guildID string, userID string, amount int, action Action) (*GuildUserScore, error) {
	return t.restClient.adjustGuildUserScore(guildID, userID, amount, action)
}

// GuildLeaderboard gets the leaderboard for a guild.
func (t *Client) GuildLeaderboard(guildID string) ([]*GuildRankedUser, error) {
	return t.restClient.guildLeaderboard(guildID)
}

// GuildUserStats gets a user's stats for a guild.
// @me is accepted for the user ID, which will retrieve stats for yourself.
func (t *Client) GuildUserStats(guildID string, userID string) (*GuildUserStats, error) {
	return t.restClient.guildUserStats(guildID, userID)
}

// User gets a Tatsumaki user profile.
func (t *Client) User(userID string) (*User, error) {
	return t.restClient.user(userID)
}
