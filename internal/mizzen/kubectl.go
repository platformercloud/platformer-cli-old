package mizzen

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// KubectlWrapper defines a wrapper around the Kubectl Binary
type KubectlWrapper struct {
	// Path to binary
	binary string
}

// NewKubectlWrapper returns a new KubectlWrapper pointer.
func NewKubectlWrapper() (*KubectlWrapper, error) {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		return nil, err
	}

	return &KubectlWrapper{
		binary: path,
	}, nil
}

// ListClusters returns the cluster list provided by kubectl config get-clusters
func (k *KubectlWrapper) ListClusters() ([]string, error) {
	b, err := k.cmd("config", "get-clusters")
	if err != nil {
		return nil, err
	}

	var clusters []string
	splitStrings := strings.Split(string(b), "\n")
	for i, c := range splitStrings {
		if i == 0 || i == len(splitStrings)-1 {
			// Ignore the first line (Header: NAME) and last (empty string)
			continue
		}
		clusters = append(clusters, strings.TrimSpace(c))
	}

	return clusters, nil
}

func (k *KubectlWrapper) cmd(args ...string) (string, error) {
	cmd := exec.Command(k.binary, args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (k *KubectlWrapper) cmdWithStdinPiped(in io.Reader, args ...string) (string, error) {
	c := exec.Command(k.binary, args...)
	c.Stdin = in

	b, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl '%s' returned an error: %s: %w", strings.Join(args, " "), string(b), err)
	}
	return string(b), nil
}
