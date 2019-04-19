package tatsumaki_go

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

func endpointGuildLeaderboard(guildId string) string {
	return endpointBaseUrl + endpointGuilds + guildId + "/leaderboard"
}

func endpointGuildUserStats(guildId string, userId string) string {
	return endpointBaseUrl + endpointGuilds + guildId + "/members/" + userId + "/stats"
}

func endpointPing() string {
	return endpointBaseUrl + "ping"
}

func endpointUser(userId string) string {
	return endpointBaseUrl + endpointUsers + userId
}

func putGuildUserPoints(guildId string, userId string) string {
	return endpointBaseUrl + endpointGuilds + guildId + "/members" + userId + "/points"
}

func putGuildUserScore(guildId string, userId string) string {
	return endpointBaseUrl + endpointGuilds + guildId + "/members" + userId + "/score"
}
