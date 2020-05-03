package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

const tokenCreateURL = "https://auth-module.dev.x.platformer.com/api/v1/serviceaccount/token/create"

// FetchPermanentToken fetches a Permanent token from the Auth API
// using the provided access token
func FetchPermanentToken(token string) (string, error) {
	b, err := json.Marshal(struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		ExpiredIn   *string `json:"expired_in"`
	}{
		// @todo - add proper values here
		"test-service account",
		"",
		nil,
	})

	req, _ := http.NewRequest("POST", tokenCreateURL, bytes.NewReader(b))
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

// SaveToken saves the given permanent token to the local config
func SaveToken(token string) {
	viper.Set("auth.token", strings.TrimSpace(token))
	viper.WriteConfig()
}

// RemoveToken removes the locally saved token
func RemoveToken() {
	viper.Set("auth.token", "")
	viper.WriteConfig()
}

// GetToken retrieves the locally stored perm.token
func GetToken() string {
	return viper.GetString("auth.token")
}
