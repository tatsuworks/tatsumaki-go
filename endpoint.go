package tatsumakigo

import "strconv"

const (
	endpointBaseURL = "https://api.tatsumaki.xyz/"

	endpointGuilds = "guilds/"
	endpointUsers  = "users/"
)

func endpointBackgroundImage(name string) string {
	return "https://www.tatsumaki.xyz/images/backgrounds/profile/" + name + ".png"
}

func endpointBadgeImage(name string) string {
	return "https://www.tatsumaki.xyz/images/badges/" + name + ".png"
}

func endpointGuildLeaderboard(guildID string, limit int) string {
	if limit != 0 {
		return endpointBaseURL + endpointGuilds + guildID + "/leaderboard?" + strconv.FormatInt(int64(limit), 10)
	}

	return endpointBaseURL + endpointGuilds + guildID + "/leaderboard"
}

func endpointGuildUserStats(guildID string, userID string) string {
	return endpointBaseURL + endpointGuilds + guildID + "/members/" + userID + "/stats"
}

func endpointPing() string {
	return endpointBaseURL + "ping"
}

func endpointUser(userID string) string {
	return endpointBaseURL + endpointUsers + userID
}

func putGuildUserPoints(guildID string, userID string) string {
	return endpointBaseURL + endpointGuilds + guildID + "/members" + userID + "/points"
}

func putGuildUserScore(guildID string, userID string) string {
	return endpointBaseURL + endpointGuilds + guildID + "/members" + userID + "/score"
}
