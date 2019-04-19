package tatsumaki_go

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type restClient struct {
	token       string
	httpClient  *http.Client
	rateLimiter *rateLimiter
}

func NewRestClient(token string) *restClient {
	// Create rest client.
	return &restClient{
		token,
		http.DefaultClient,
		newRateLimiter(),
	}
}

func (rc *restClient) User(userId string) (*User, error) {
	// Create request.
	req, err := http.NewRequest("GET", endpointUser(userId), nil)
	if err != nil {
		return nil, errorRequestFailed(err)
	}

	// Set headers.
	rc.setHeaders(req)

	// Wait for rate limit clearance.
	rc.rateLimiter.Lock()
	rc.rateLimiter.acquire()

	// Execute request.
	response, err := rc.httpClient.Do(req)

	// Store last request time and unlock rate limiter.
	rc.rateLimiter.lastRequest = time.Now()
	rc.rateLimiter.Unlock()

	if err != nil {
		return nil, errorResponseFailed(err)
	}

	// Check if response was successful.
	if response.StatusCode != 200 {
		// Attempt to parse error JSON.
		var tatsuErr map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&tatsuErr)
		if err != nil {
			return nil, errorResponseFailed(nil)
		}
		return nil, errorResponseFailed(errors.New(tatsuErr["message"].(string)))
	}

	// Parse response.
	var jsonResponse map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	// Check if there was an error decoding.
	if err != nil {
		return nil, errorResponseFailed(err)
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

func (rc *restClient) setHeaders(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", rc.token)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "Tatsumaki Go/1.0.0-alpha (Hassie, https://github.com/hassieswift621/tatsumaki-go")
}
