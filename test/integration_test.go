package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func login(user map[string]any) (io.Reader, error) {
	return mapToBody(map[string]any{
		"email":    user["email"],
		"password": user["password"],
	})
}

// helper function to convert the response body to a map
func bodyToMap(resp *http.Response) (map[string]any, error) {
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var structuredResponse map[string]any
	json.Unmarshal(content, &structuredResponse)
	return structuredResponse, nil
}

func mapToBody(m map[string]any) (io.Reader, error) {
	content, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(content), nil
}

func TestAPI(t *testing.T) {
	assert := assert.New(t)
	client := http.Client{}
	// create 20 users
	users := []map[string]any{}
	for range 2 {
		resp, err := client.Post("http://localhost:5000/user/create", "application/json", nil)
		assert.Equal(http.StatusCreated, resp.StatusCode)
		assert.NoError(err)
		userResponse, err := bodyToMap(resp)
		assert.NoError(err)

		assert.Contains(userResponse, "result")
		assert.Contains(userResponse["result"], "id")
		assert.IsType(map[string]any{}, userResponse["result"])
		user := userResponse["result"].(map[string]any)
		assert.NotNil(user["id"])
		assert.Contains(user, "email")
		assert.NotNil(user["email"])
		assert.Contains(user, "password")
		assert.NotNil(user["password"])

		if len(users) != 2 {
			users = append(users, user)
		}
	}

	// get first user
	payloadLoginUser1, err := login(users[0])
	resp, err := client.Post("http://localhost:5000/login", "application/json", payloadLoginUser1)
	assert.Equal(http.StatusOK, resp.StatusCode)

	tokenResponseUser1, err := bodyToMap(resp)
	assert.NoError(err)

	assert.Contains(tokenResponseUser1, "result")
	assert.Contains(tokenResponseUser1["result"], "token")
	assert.NotNil(tokenResponseUser1["result"].(map[string]any)["token"])
	tokenUser1 := tokenResponseUser1["result"].(map[string]any)["token"].(string)

	discoverReqUser1, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/discover", nil)
	discoverReqUser1.Header.Set(
		"Authorization",
		tokenUser1,
	)
	resp, err = client.Do(discoverReqUser1)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	discoverResponseUser1, err := bodyToMap(resp)
	assert.NoError(err)
	assert.IsType([]any{}, discoverResponseUser1["results"])
	discoverResultsUser1 := discoverResponseUser1["results"].([]any)
	assert.Greater(len(discoverResultsUser1), 1)
	for _, result := range discoverResultsUser1 {
		assert.IsType(map[string]any{}, result)
		assert.Contains(result, "id")
		assert.NotEqual(result.(map[string]any)["id"], users[0]["id"])
	}

	// randomly pick a user
	matchUser := users[1]

	// now lets swipe date
	payloadSwipeUser1, err := mapToBody(map[string]any{
		"userID":     matchUser["id"],
		"preference": "YES",
	})
	assert.NoError(err)
	swipeReqUser1, _ := http.NewRequest(
		http.MethodPost,
		"http://localhost:5000/swipe",
		payloadSwipeUser1,
	)
	swipeReqUser1.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	swipeReqUser1.Header.Set("Authorization", tokenUser1)

	resp, err = client.Do(swipeReqUser1)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	// check user does not exist in the list of users
	firstSwipeResponse, err := bodyToMap(resp)
	assert.NoError(err)
	assert.Contains(firstSwipeResponse, "result")
	assert.IsType(map[string]any{}, firstSwipeResponse["result"])
	firstSwipeResult := firstSwipeResponse["result"].(map[string]any)
	assert.Contains(firstSwipeResult, "matched")
	assert.NotContains(firstSwipeResult, "matchID")

	// todo: now check discover to ensure we don't have the match anymore
	// now lets switch user ot matched and swipe as well so we see they are matched
	payloadLoginUser2, err := login(users[1])
	assert.NoError(err)

	resp, err = client.Post("http://localhost:5000/login", "application/json", payloadLoginUser2)
	assert.Equal(http.StatusOK, resp.StatusCode)

	tokenResponseUser2, err := bodyToMap(resp)
	assert.NoError(err)

	assert.Contains(tokenResponseUser2, "result")
	assert.Contains(tokenResponseUser2["result"], "token")
	assert.NotNil(tokenResponseUser2["result"].(map[string]any)["token"])
	tokenUser2 := tokenResponseUser2["result"].(map[string]any)["token"].(string)

	discoverReqUser2, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/discover", nil)
	discoverReqUser2.Header.Set(
		"Authorization",
		tokenUser2,
	)
	resp, err = client.Do(discoverReqUser2)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	discoverResponseUser2, err := bodyToMap(resp)
	assert.NoError(err)
	assert.IsType([]any{}, discoverResponseUser2["results"])
	discoverResultsUser2 := discoverResponseUser2["results"].([]any)
	assert.Greater(len(discoverResultsUser2), 1)
	var discoverUser2 map[string]any
	for _, result := range discoverResultsUser2 {
		if result.(map[string]any)["id"] == users[0]["id"] {
			discoverUser2 = result.(map[string]any)
		}
	}
	assert.NotNil(discoverUser2)

	// now lets swipe date
	payloadSwipeUser2, err := mapToBody(map[string]any{
		"userID":     matchUser["id"],
		"preference": "YES",
	})
	assert.NoError(err)
	swipeReqUser2, _ := http.NewRequest(
		http.MethodPost,
		"http://localhost:5000/swipe",
		payloadSwipeUser2,
	)
	swipeReqUser2.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	// todo: this is incorret token
	swipeReqUser2.Header.Set("Authorization", tokenUser2)

	resp, err = client.Do(swipeReqUser2)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	// check user does not exist in the list of users
	secondSwipeResponse, err := bodyToMap(resp)
	assert.NoError(err)
	assert.Contains(secondSwipeResponse, "result")
	assert.IsType(map[string]any{}, secondSwipeResponse["result"])
	secondSwipeResult := secondSwipeResponse["result"].(map[string]any)
	assert.Contains(secondSwipeResult, "matched")
	assert.Contains(secondSwipeResult, "matchID")
}
