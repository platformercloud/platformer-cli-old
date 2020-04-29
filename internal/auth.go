package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

type permanentTokenRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ExpiredIn   *string   `json:"expired_in"`
}

// Create a Permanent token
func CreatePermanentToken(token string) (string, error) {
	client := &http.Client{}
	URL := "https://auth-module.dev.x.platformer.com/api/v1/serviceaccount/token/create"
	request := permanentTokenRequestBody{
		Name:        "test-service account",
		Description: "",
		ExpiredIn:   nil,
	}

	byteVal, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body : %w", err)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewReader(byteVal))
	if err != nil {
		return "", fmt.Errorf("error creating new request : %w", err)
	}

	req.Header.Set("Authorization", strings.TrimSpace(token))
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error reading response body : %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body : %w", err)
	}

	responseString := string(body)

	if !gjson.Get(responseString, "success").Bool() {
		return "", fmt.Errorf("error getting permanent token : %w", err)
	}

	permanentToken := gjson.Get(responseString, "data.token")
	return permanentToken.Str, nil
}
