package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func saveGlobalFile(filePath string, value string) (bool, error) {
	var dir string
	dirPath := "/.platformer"
	dir, err := GetOSRootDir()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(dir + filePath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir+dirPath, 0755)
		if errDir != nil {
			return false, errDir
		}

	}

	f, err := os.Create(dir + filePath)
	if err != nil {
		return false, fmt.Errorf("error file creating. %s", err)

	}
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(value)
	if err != nil {
		return false, fmt.Errorf("error token wrting %s", err)
	}
	_ = writer.Flush()

	return true, nil
}

func GetLocallyStoredToken() (string, error) {

	var dir string
	dir, err := GetOSRootDir()
	if err != nil {
		fmt.Println(err)
	}

	token, err := ioutil.ReadFile(dir + "/.platformer/token")
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	return string(token), nil
}

func SaveOrganizationGlobally(organizationID string) (bool, error) {
	return saveGlobalFile("/.platformer/organizations", organizationID)
}

func SaveProjectGlobally(projectID string) (bool, error) {
	return saveGlobalFile("/.platformer/projects", projectID)
}
