package auth

import (
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/util"
	"net/http"

	"github.com/platformercloud/platformer-cli/internal/config"
)

// IsLoggedIn checks if the user is logged in with the CLI
// and if the saved permanent token is valid
func IsLoggedIn() bool {
	t := config.GetToken()
	if t == "" {
		return false
	}
	req, _ := http.NewRequest("PUT", util.AuthValidTokenURL, nil)
	req.Header.Set("Authorization", t)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[internal] Error validating token: %v\n", err)
		config.RemoveToken()
		return false
	}
	if resp.StatusCode != 200 {
		config.RemoveToken()
		return false
	}
	return true
}
