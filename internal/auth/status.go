package auth

import "github.com/spf13/viper"

// IsLoggedIn checks if the user is logged in with the CLI
// and if the saved permanent token is valid
func IsLoggedIn() bool {
	// @todo - validate this token with a new endpoint from achala
	// check if the token is still valid
	t := viper.GetString("auth.token")
	if t == "" {
		return false
	}
	return true
}
