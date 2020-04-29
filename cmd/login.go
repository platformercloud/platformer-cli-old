package cmd

import (
	"bufio"
	"fmt"
	"gitlab.platformer.com/chamod.p/platformer/internal"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	conf *oauth2.Config
)

func login() {
	conf = &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://console.dev.x.platformer.com/cli-login",
		},
		// CLI callback URL
		RedirectURL: "http://localhost:9999",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	loginUrl := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)

	if err := open.Run(loginUrl); err != nil {
		log.Fatalf("cannot open browser: %s", err)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("Authentication URL: %s\n", loginUrl)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD, CONNECT")
			w.Header().Add("Access-Control-Allow-Origins", "*")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Token")

			w.WriteHeader(http.StatusOK)
			return
		}

		xToken := r.Header.Get("x-token")
		permanentToken, err := internal.CreatePermanentToken(xToken)
		if err != nil {
			log.Fatalf("error creating permanent token %s", err)
		}

		err = saveToken(permanentToken)
		if err != nil {
			w.WriteHeader(400)
			fmt.Printf("error Message: %s", err)
			os.Exit(1)
		}

		_, err = w.Write([]byte("Success"))
		if err != nil {
			log.Fatalf("response error %s", err)
		}

		w.WriteHeader(200)

		fmt.Println(color.GreenString("Successfully logged"))
		os.Exit(0)
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

	server := http.ListenAndServe(":9999", c.Handler(mux))

	log.Fatal(server)
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
	Short: "log the CLI into Platformer Cloud",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
