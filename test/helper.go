package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func loginEndpoint(assert *assert.Assertions, user map[string]any) string {
	client := http.Client{}
	payload, err := mapToBody(map[string]any{
		"email":    user["email"],
		"password": user["password"],
	})
	assert.NoError(err)

	resp, err := client.Post("http://localhost:5000/login", "application/json", payload)
	assert.Equal(http.StatusOK, resp.StatusCode)

	tokenResponse, err := bodyToMap(resp)
	assert.NoError(err)

	assert.Contains(tokenResponse, "result")
	assert.Contains(tokenResponse["result"], "token")
	assert.NotNil(tokenResponse["result"].(map[string]any)["token"])
	token := tokenResponse["result"].(map[string]any)["token"].(string)
	return token
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

func findUser(users []any, user map[string]any) map[string]any {
	var foundUser map[string]any
	for _, result := range users {
		if result.(map[string]any)["id"] == user["id"] {
			foundUser = result.(map[string]any)
		}
	}
	return foundUser
}

func discoverEndpoint(assert *assert.Assertions, token string) []any {
	client := http.Client{}
	discoverReq, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/discover", nil)
	discoverReq.Header.Set(
		"Authorization",
		token,
	)
	resp, err := client.Do(discoverReq)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	discoverResponse, err := bodyToMap(resp)
	assert.NoError(err)
	assert.IsType([]any{}, discoverResponse["results"])
	discoverResultsUser1 := discoverResponse["results"].([]any)
	assert.Greater(len(discoverResultsUser1), 1)
	return discoverResultsUser1
}

func swipeEndpoint(assert *assert.Assertions, token string, body map[string]any) map[string]any {
	client := http.Client{}
	payload, err := mapToBody(body)
	assert.NoError(err)
	swipeReq, _ := http.NewRequest(
		http.MethodPost,
		"http://localhost:5000/swipe",
		payload,
	)
	swipeReq.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	swipeReq.Header.Set("Authorization", token)

	resp, err := client.Do(swipeReq)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	firstSwipeResponse, err := bodyToMap(resp)
	assert.NoError(err)
	assert.Contains(firstSwipeResponse, "result")
	assert.IsType(map[string]any{}, firstSwipeResponse["result"])
	firstSwipeResult := firstSwipeResponse["result"].(map[string]any)
	return firstSwipeResult
}

func createUserEndpoint(assert *assert.Assertions) map[string]any {
	client := http.Client{}
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
	return user
}
