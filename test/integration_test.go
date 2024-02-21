package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	resp, err := client.Post("http://localhost:5000/user/create", "application/json", nil)
	assert.NoError(err)

	structuredResponse, err := bodyToMap(resp)
	assert.NoError(err)

	assert.Contains(structuredResponse, "result")
	assert.Contains(structuredResponse["result"], "email")
	assert.NotNil(structuredResponse["result"].(map[string]any)["email"])
	assert.Contains(structuredResponse["result"], "password")
	assert.NotNil(structuredResponse["result"].(map[string]any)["password"])

	payload, err := mapToBody(map[string]any{
		"email":    structuredResponse["result"].(map[string]any)["email"],
		"password": structuredResponse["result"].(map[string]any)["password"],
	})
	resp, err = client.Post("http://localhost:5000/login", "application/json", payload)

	structuredResponse, err = bodyToMap(resp)
	assert.NoError(err)

	assert.Contains(structuredResponse, "result")
	assert.Contains(structuredResponse["result"], "token")
	assert.NotNil(structuredResponse["result"].(map[string]any)["token"])
}
