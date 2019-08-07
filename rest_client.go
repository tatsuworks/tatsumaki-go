package tatsumakigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// RestClient is the low level client which handles the requests.
type restClient struct {
	token       string
	httpClient  *http.Client
	rateLimiter *rateLimiter
}

func newRestClient(token string) *restClient {
	// Create rest client.
	return &restClient{
		token,
		http.DefaultClient,
		newRateLimiter(),
	}
}

func (rc *restClient) adjustGuildUserPoints(guildId string, userId string, amount int, action Action) (*GuildUserPoints, error) {
	// Check if amount is 0 and action is add or remove.
	if amount == 0 && (action == ActionAdd || action == ActionRemove) {
		return nil, errorAdjustInvalid()
	}

	// Check if amount is between 0 and 50,000 (inclusive).
	if amount < 0 || amount > 50000 {
		return nil, errorAdjustBounds()
	}

	// Make request.
	body, err := rc.makePutRequest(putGuildUserPoints(guildId, userId),
		adjustGuildUserPoints{amount, string(action)})
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var guildUserPoints GuildUserPoints
	err = json.NewDecoder(body).Decode(&guildUserPoints)
	if err != nil {
		return nil, errorParseFailed(err)
	}

	return &guildUserPoints, nil
}

func (rc *restClient) adjustGuildUserScore(guildId string, userId string, amount int, action Action) (*GuildUserScore, error) {
	// Check if amount is 0 and action is add or remove.
	if amount == 0 && (action == ActionAdd || action == ActionRemove) {
		return nil, errorAdjustInvalid()
	}

	// Check if amount is between 0 and 50,000 (inclusive).
	if amount < 0 || amount > 50000 {
		return nil, errorAdjustBounds()
	}

	// Make request.
	body, err := rc.makePutRequest(putGuildUserScore(guildId, userId),
		adjustGuildUserScore{amount, string(action)})
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var guildUserScore GuildUserScore
	err = json.NewDecoder(body).Decode(&guildUserScore)
	if err != nil {
		return nil, errorParseFailed(err)
	}

	return &guildUserScore, nil
}

func (rc *restClient) guildLeaderboard(guildId string) ([]*GuildRankedUser, error) {
	// Make request.
	body, err := rc.makeGetRequest(endpointGuildLeaderboard(guildId))
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var guildLeaderboard []*GuildRankedUser
	err = json.NewDecoder(body).Decode(&guildLeaderboard)

	// Check if there was an error decoding.
	if err != nil {
		return nil, errorParseFailed(err)
	}

	// Iterate over the array and remove the null elements received from the response.
	// Once we find the first null element, we simply cut the array at that index.
	for i, v := range guildLeaderboard {
		if v == nil {
			guildLeaderboard = guildLeaderboard[:i]
			break
		}
	}

	return guildLeaderboard, nil
}

func (rc *restClient) guildUserStats(guildId string, userId string) (*GuildUserStats, error) {
	// Make request.
	body, err := rc.makeGetRequest(endpointGuildUserStats(guildId, userId))
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var guildUserStats GuildUserStats
	err = json.NewDecoder(body).Decode(&guildUserStats)

	// Check if there was an error decoding.
	if err != nil {
		return nil, errorParseFailed(err)
	}

	return &guildUserStats, nil
}

func (rc *restClient) ping() (*Ping, error) {
	// Make request.
	body, err := rc.makeGetRequest(endpointPing())
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var ping Ping
	err = json.NewDecoder(body).Decode(&ping)

	// Check if there was an error decoding.
	if err != nil {
		return nil, errorParseFailed(err)
	}

	return &ping, nil
}

func (rc *restClient) user(userId string) (*User, error) {
	// Make request.
	body, err := rc.makeGetRequest(endpointUser(userId))
	if err != nil {
		return nil, err
	}

	// Defer closing body.
	defer body.Close()

	// Parse response.
	var jsonResponse map[string]interface{}
	err = json.NewDecoder(body).Decode(&jsonResponse)

	// Check if there was an error decoding.
	if err != nil {
		return nil, errorParseFailed(err)
	}

	// Extract JSON.
	var user User
	user.AvatarURL = jsonResponse["avatar_url"].(string)
	user.Background = &Background{
		endpointBackgroundImage(jsonResponse["background"].(string)),
		jsonResponse["background"].(string),
	}
	badgeSlots := jsonResponse["badgeSlots"].([]interface{})
	for i, badgeName := range badgeSlots {
		if badgeName == nil {
			user.BadgeSlots = append(user.BadgeSlots, &BadgeSlot{
				nil,
				i + 1,
			})
		} else {
			user.BadgeSlots = append(user.BadgeSlots, &BadgeSlot{
				&Badge{
					endpointBadgeImage(badgeName.(string)),
					badgeName.(string),
				},
				i + 1,
			})
		}
	}
	user.Credits = int64(jsonResponse["credits"].(float64))
	user.InfoBox = jsonResponse["info_box"].(string)
	user.Level = int64(jsonResponse["level"].(float64))
	levelProgress := jsonResponse["xp"].([]interface{})
	user.LevelProgress = &LevelProgress{
		int64(levelProgress[0].(float64)),
		int64(levelProgress[1].(float64)),
	}
	user.Name = jsonResponse["name"].(string)
	user.Rank = int64(jsonResponse["rank"].(float64))
	user.Reputation = int64(jsonResponse["reputation"].(float64))
	user.Title = jsonResponse["title"].(string)
	user.TotalXp = int64(jsonResponse["total_xp"].(float64))

	return &user, nil
}

// MakeGetRequest makes a GET request to the API.
func (rc *restClient) makeGetRequest(endpoint string) (io.ReadCloser, error) {
	// Create request.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errorRequestFailed(err)
	}

	// Set headers.
	rc.setHeaders(req)

	// Wait for rate limit clearance.
	rc.rateLimiter.Lock()
	rc.rateLimiter.acquire()

	// Execute request.
	res, err := rc.httpClient.Do(req)

	// Store last request time and unlock rate limiter.
	rc.rateLimiter.lastRequest = time.Now()
	rc.rateLimiter.Unlock()

	if err != nil {
		return nil, errorResponseFailed(err)
	}

	// Check if res was successful.
	if res.StatusCode != 200 {
		// Attempt to parse error JSON.
		var tatsuErr tatsumakiError
		err := json.NewDecoder(res.Body).Decode(&tatsuErr)
		if err != nil {
			return nil, errorResponseFailed(nil)
		}
		res.Body.Close()
		return nil, errorResponseFailed(errors.New(tatsuErr.Message))
	}

	return res.Body, nil
}

// MakePutRequest makes a PUT request to the API.
func (rc *restClient) makePutRequest(endpoint string, body interface{}) (io.ReadCloser, error) {
	// Encode body into JSON.
	encoded, err := json.Marshal(&body)
	if err != nil {
		return nil, errorRequestFailed(err)
	}

	// Create request.
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(encoded))
	if err != nil {
		return nil, errorRequestFailed(err)
	}

	// Set headers.
	rc.setHeaders(req)

	// Wait for rate limit clearance.
	rc.rateLimiter.Lock()
	rc.rateLimiter.acquire()

	// Execute request.
	res, err := rc.httpClient.Do(req)

	// Store last request time and unlock rate limiter.
	rc.rateLimiter.lastRequest = time.Now()
	rc.rateLimiter.Unlock()

	// Check if there is an error with the response.
	if err != nil {
		return nil, errorResponseFailed(err)
	}

	// Check if response was successful.
	if res.StatusCode != 200 {
		// Attempt to parse error JSON.
		var tatsuErr tatsumakiError
		err := json.NewDecoder(res.Body).Decode(&tatsuErr)
		if err != nil {
			return nil, errorResponseFailed(nil)
		}
		res.Body.Close()
		return nil, errorResponseFailed(errors.New(tatsuErr.Message))
	}

	return res.Body, nil
}

// SetHeaders sets the headers for each request.
func (rc *restClient) setHeaders(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", rc.token)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "Tatsumaki Go/1.1.0-release (Hassie, https://github.com/hassieswift621/tatsumaki-go)")
}
