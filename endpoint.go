package tatsumakigo

const (
	endpointBaseUrl = "https://api.tatsumaki.xyz/"

	endpointGuilds = "guilds/"
	endpointUsers  = "users/"
)

func endpointBackgroundImage(name string) string {
	return "https://www.tatsumaki.xyz/images/backgrounds/profile/" + name + ".png"
}

func endpointBadgeImage(name string) string {
	return "https://www.tatsumaki.xyz/images/badges/" + name + ".png"
}

func endpointGuildLeaderboard(guildID string) string {
	return endpointBaseUrl + endpointGuilds + guildID + "/leaderboard"
}

func endpointGuildUserStats(guildID string, userID string) string {
	return endpointBaseUrl + endpointGuilds + guildID + "/members/" + userID + "/stats"
}

func endpointPing() string {
	return endpointBaseUrl + "ping"
}

func endpointUser(userID string) string {
	return endpointBaseUrl + endpointUsers + userID
}

func putGuildUserPoints(guildID string, userID string) string {
	return endpointBaseUrl + endpointGuilds + guildID + "/members" + userID + "/points"
}

func putGuildUserScore(guildID string, userID string) string {
	return endpointBaseUrl + endpointGuilds + guildID + "/members" + userID + "/score"
}
