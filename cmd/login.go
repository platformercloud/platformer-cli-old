package cmd

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"

	"github.com/fatih/color"
	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	oauthConfig *oauth2.Config
)

const (
	authURL     string = "https://console.dev.x.platformer.com/cli-login"
	port        string = ":9999"
	redirectURL string = "http://127.0.0.1" + port
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into Platformer through the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		HandleErrorAndExit(login())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login() error {
	if auth.IsLoggedIn() {
		return &UserError{fmt.Errorf("you are already logged in")}
	}

	server := &http.Server{Addr: port}
	done := make(chan string)
	errc := make(chan error)

	// Start the server on a seperate go routine
	go startServerAndAwaitToken(server, done, errc)

	oauthConfig = &oauth2.Config{
		Endpoint:    oauth2.Endpoint{AuthURL: authURL},
		RedirectURL: redirectURL,
	}

	loginURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit this URL on your device to log in:\n%s\n", loginURL)
	fmt.Println(color.CyanString("\nYou will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)

	// Redirect user to CLI login page
	if err := open.Run(loginURL); err != nil {
		return &UserError{fmt.Errorf("cannot open browser: %s", err)}
	}

	// Block until a response from the server is recieved
	// or until it times out.
	select {
	case token := <-done:
		server.Close()
		permanentToken, err := auth.FetchPermanentToken(token)
		if err != nil {
			return &InternalError{err, "failed to sign in"}
		}

		config.SaveToken(permanentToken)
		fmt.Println(color.GreenString("Successfully logged in!"))
		return nil

	case err := <-errc:
		server.Close()
		return &InternalError{err, "cannot listen on port " + port}

	case <-time.After(2 * time.Minute):
		server.Close()
		return &UserError{fmt.Errorf("timed out, try again")}
	}
}

func startServerAndAwaitToken(server *http.Server, tokenChan chan<- string, errc chan<- error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-token")
		if token == "" {
			errc <- &UserError{fmt.Errorf("failed to log in")}
			w.WriteHeader(400)
			return
		}
		tokenChan <- token
		w.WriteHeader(200)
	})

	c := cors.New(cors.Options{
		// @TODO: add production/staging URLs
		AllowedMethods:     []string{http.MethodPost, http.MethodOptions, http.MethodConnect},
		AllowedOrigins:     []string{"http://localhost:3000", "https://console.dev.x.platformer.com", "http://localhost:9999"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		AllowCredentials:   false,
		OptionsPassthrough: false,
		MaxAge:             120,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	server.Handler = c.Handler(mux)
	if err := server.ListenAndServe(); err != nil {
		errc <- err
	}
}
