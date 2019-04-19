package tatsumakigo

// TatsumakiClient is the main interface to interact with the API.
type TatsumakiClient struct {
	rc *restClient
}

// New creates a new instance of the Tatsumaki client.
func New(token string) *TatsumakiClient {
	return &TatsumakiClient{
		newRestClient(token),
	}
}

// GuildLeaderboard gets the leaderboard for a guild.
func (c *TatsumakiClient) GuildLeaderboard(guildId string) ([]*GuildRankedUser, error) {
	return c.rc.guildLeaderboard(guildId)
}

// GuildUserStats gets a user's stats for a guild.
// @me is accepted for the user ID, which will retrieve stats for yourself.
func (c *TatsumakiClient) GuildUserStats(guildId string, userId string) (*GuildUserStats, error) {
	return c.rc.guildUserStats(guildId, userId)
}

// User gets a Tatsumaki user's profile.
func (c *TatsumakiClient) User(userId string) (*User, error) {
	return c.rc.user(userId)
}
