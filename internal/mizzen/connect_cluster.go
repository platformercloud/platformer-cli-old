package mizzen

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

const (
	mizzenAPI         = "https://mizzen.dev.x.platformer.com"
	registrationURL   = mizzenAPI + "/api/v1/cluster"
	yamlGenerationURL = mizzenAPI + "/generate/v1/agent/"
)

type credentials struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ConnectAndInstallAgent registers the cluster with the Mizzen API and installs the in-cluster mx-agent
func ConnectAndInstallAgent(kw *KubectlWrapper, orgID string, projectID string, clusterName string) error {
	credentials, err := register(orgID, projectID, clusterName)
	if err != nil {
		return &cli.InternalError{
			Message: "Failed to register cluster",
			Err:     err,
		}
	}

	if err := installAgent(kw, credentials); err != nil {
		return &cli.InternalError{
			Message: "Failed to install in-cluster agent. Please check if kubectl has access to the requested cluster",
			Err:     err,
		}
	}

	return nil
}

func register(orgID string, projectID string, clusterName string) (*credentials, error) {
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(struct {
		ClusterType    string   `json:"cluster_type"`
		ClusterName    string   `json:"cluster_name"`
		ProjectID      string   `json:"project_id"`
		OrganizationID string   `json:"organization_id"`
		WhitelistIPs   []string `json:"whitelist_ips"`
	}{
		"private",
		clusterName,
		projectID,
		orgID,
		[]string{}, // Whitelist IPs are not set using the CLI
	})

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	r, err := client.Post(registrationURL, "application/json", &body)
	if err != nil {
		return nil, fmt.Errorf("api request failed (register cluster): %w", err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if r.StatusCode >= 400 {
		return nil, fmt.Errorf("bad request error: %s", string(b))
	}

	var creds credentials
	if err := json.Unmarshal(b, &creds); err != nil {
		return nil, fmt.Errorf("invalid response (register cluster): %w", err)
	}

	return &creds, nil
}

func installAgent(kw *KubectlWrapper, creds *credentials) error {
	encodedToken := base64.StdEncoding.EncodeToString([]byte(creds.ClientID + ";" + creds.ClientSecret))
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	r, err := client.Get(yamlGenerationURL + encodedToken)
	if err != nil {
		return fmt.Errorf("failed to get agent installation yaml: %w", err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if _, err := kw.cmdWithStdinPiped(bytes.NewBuffer(b), "apply", "-f", "-"); err != nil {
		return fmt.Errorf("failed to install agent: %w", err)
	}

	return nil
}
