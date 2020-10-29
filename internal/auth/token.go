package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/util"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// FetchPermanentToken fetches a Permanent token from the Auth API
// using the provided access token
func FetchPermanentToken(token string) (string, error) {
	b, err := json.Marshal(struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		ExpiredIn   *string `json:"expired_in"`
	}{
		"cli-service account",
		"Getting token for CLI use",
		nil,
	})

	req, _ := http.NewRequest("POST", util.AuthTokenCreateURL, bytes.NewReader(b))
	req.Header.Set("Authorization", strings.TrimSpace(token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	responseString := string(body)
	if !gjson.Get(responseString, "success").Bool() {
		return "", fmt.Errorf("auth API failed to return a permanent token: %w", err)
	}

	permanentToken := gjson.Get(responseString, "data.token")
	return permanentToken.Str, nil
}
