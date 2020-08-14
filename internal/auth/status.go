package auth

import (
	"fmt"
	"net/http"

	"github.com/platformercloud/platformer-cli/internal/config"
)

const (
	validateTokenURL = "https://api.ambassador.dev.platformer.com/auth/api/v1/user/logintime"
)

// IsLoggedIn checks if the user is logged in with the CLI
// and if the saved permanent token is valid
func IsLoggedIn() bool {
	t := config.GetToken()
	if t == "" {
		return false
	}
	req, _ := http.NewRequest("PUT", validateTokenURL, nil)
	req.Header.Set("Authorization", t)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[internal] Error validating token: %v", err)
		config.RemoveToken()
		return false
	}
	if resp.StatusCode != 200 {
		config.RemoveToken()
		return false
	}
	return true
}
