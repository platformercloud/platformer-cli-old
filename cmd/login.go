package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gitlab.platformer.com/project-x/platformer-cli/internal"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"

	"github.com/fatih/color"
	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	conf *oauth2.Config
)

const port string = ":9999"

func login() error {
	conf = &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://console.dev.x.platformer.com/cli-login",
		},
		RedirectURL: "http://127.0.0.1" + port,
	}

	loginURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit this URL on your device to log in:\n%s\n", loginURL)
	fmt.Println(color.CyanString("\nYou will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)

	// Redirect user to consent page to ask for permission
	if err := open.Run(loginURL); err != nil {
		return UserError{fmt.Errorf("cannot open browser: %s", err)}
	}

	// Start the server and wait for user to log in.
	server := &http.Server{Addr: port}
	tokenChan := make(chan string)
	go startServerAndAwaitToken(server, tokenChan)
	token := <-tokenChan
	server.Close()

	permanentToken, err := auth.CreatePermanentToken(token)
	if err != nil {
		return InternalError{err, "failed to sign in"}
	}

	err = saveToken(permanentToken)
	if err != nil {
		return err
	}

	return nil
}

func startServerAndAwaitToken(server *http.Server, tokenChan chan<- string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-token")
		tokenChan <- token
		w.WriteHeader(200)
	})

	c := cors.New(cors.Options{
		// @TODO: add production/staging URLs
		AllowedMethods:     []string{http.MethodPost, http.MethodOptions, http.MethodConnect},
		AllowedOrigins:     []string{"http://localhost:3000", "https://console.dev.x.platformer.com", "http://localhost:9999"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		MaxAge:             120,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	server.Handler = c.Handler(mux)
	return server.ListenAndServe()
}

// create .platformer and store the token
func fileCreate(dir string, token string) {

	// create .platformer dir
	_, err := os.Stat(dir + "/.platformer/token")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir+"/.platformer/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	f, err := os.Create(dir + "/.platformer/token")
	if err != nil {
		log.Fatal("error file creating. ", err)

	}
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(token)
	if err != nil {
		log.Fatalf("error token wrting %s", err)
	}
	_ = writer.Flush()
}

// Save permanent token in local
func saveToken(token string) error {

	var dir string
	dir, err := internal.GetOSRootDir()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fileCreate(dir, token)

	// Validate created token file
	_, err = os.Stat(dir + "/.platformer/token")
	os.IsNotExist(err)
	// TOKEN file does not exist
	return err
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into Platformer through the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
