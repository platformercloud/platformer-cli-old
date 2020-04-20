package tokens

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Generate a token
func createToken() string {
	resp, err := http.PostForm("https://auth-module.dev.x.platformer.com/api/v1/user/login",
		url.Values{"email": {"diaspositive@gmail.com"}, "password": {"12345678"}})
	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	tokenValue := gjson.Get(responseString, "data.id_token")

	return tokenValue.Str
}


//mux := http.NewServeMux()
//mux.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
//	xToken := r.Header.Get("x-token")
//	_, _ = w.Write([]byte("Success"))
//	w.WriteHeader(200)
//
//	err := saveToken(internal.CreatePermanentToken(xToken))
//	if err != nil {
//		// deviate from happy path
//		w.WriteHeader(400)
//		fmt.Println("Error Message: %w", err)
//		os.Exit(1)
//	}
//
//	w.WriteHeader(200)
//	fmt.Println(color.GreenString("Successfully logged"))
//	os.Exit(0)
//
//})