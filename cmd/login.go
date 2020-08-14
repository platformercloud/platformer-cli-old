package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/platformer-com/platformer-cli/internal/auth"
	"github.com/platformer-com/platformer-cli/internal/cli"
	"github.com/platformer-com/platformer-cli/internal/config"

	"github.com/gookit/color"
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
		cli.HandleErrorAndExit(login())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login() error {
	if auth.IsLoggedIn() {
		return &cli.UserError{Message: "you are already logged in"}
	}

	server := &http.Server{Addr: port}
	done := make(chan string)
	errc := make(chan error)

	// Start the server on a separate go routine
	go startServerAndAwaitToken(server, done, errc)

	oauthConfig = &oauth2.Config{
		Endpoint:    oauth2.Endpoint{AuthURL: authURL},
		RedirectURL: redirectURL,
	}

	loginURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit this URL on your device to log in:\n%s\n", loginURL)
	color.FgCyan.Println("\nYou will now be taken to your browser for authentication")
	time.Sleep(1 * time.Second)

	// Redirect user to CLI login page
	if err := open.Run(loginURL); err != nil {
		return &cli.UserError{Message: fmt.Sprintf("cannot open browser: %s", err)}
	}

	// Block until a response from the server is received
	// or until it times out.
	select {
	case token := <-done:
		_ = server.Close()
		permanentToken, err := auth.FetchPermanentToken(token)
		if err != nil {
			return &cli.InternalError{Err: err, Message: "failed to sign in"}
		}

		config.SaveToken(permanentToken)
		color.FgGreen.Println("Successfully logged in!")
		return nil

	case err := <-errc:
		_ = server.Close()
		return &cli.InternalError{Err: err, Message: "cannot listen on port " + port}

	case <-time.After(2 * time.Minute):
		_ = server.Close()
		return &cli.UserError{Message: "timed out, try again"}
	}
}

func startServerAndAwaitToken(server *http.Server, tokenChan chan<- string, errc chan<- error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-token")
		if token == "" {
			errc <- &cli.UserError{Message: "failed to log in"}
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
