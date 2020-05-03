package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

type projectsResponse struct {
	Success bool      `json:"success"`
	Data    []Project `json:"data"`
	Message string    `json:"message"`
}

type Project struct {
	ProjectID      string `json:"projectId"`
	ProjectName    string `json:"projectName"`
	Roles          []role `json:"roles"`
	IsProjectOwner bool   `json:"isProjectOwner"`
}

type role struct {
	CollectionID string   `json:"collection_id"`
	RolesID      []string `json:"roles_id"`
}

func getOrganizationID() (string, error) {
	var dir string
	dir, err := GetOSRootDir()
	if err != nil {
		return "", fmt.Errorf("error getting root dir %s", err)
	}

	organizationID, err := ioutil.ReadFile(dir + "/.platformer/organizations")
	if err != nil {
		return "", fmt.Errorf("error getting organization locally stored token %s", err)
	}

	return string(organizationID), nil
}

func GetProjects() ([]Project, error) {

	client := &http.Client{}
	organizationID, err := getOrganizationID()
	if err != nil {
		return nil, fmt.Errorf("could not get orgnization id %s", err)
	}

	url := "https://auth-module.dev.x.platformer.com/api/v1/organization/project/list/" + organizationID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request : %w", err)
	}

	token := config.GetToken()
	// if err != nil {
	// 	return nil, fmt.Errorf("error getting token %s", err)
	// }

	req.Header.Add("Authorization", strings.TrimSpace(token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body : %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body : %w", err)
	}

	responseString := string(body)

	var projectListResponse projectsResponse
	err = json.Unmarshal([]byte(responseString), &projectListResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json %s", err)
	}

	return projectListResponse.Data, nil
}

func GetProjectsNames(projectList []Project) []string {
	var projectNames []string

	for _, project := range projectList {
		projectNames = append(projectNames, project.ProjectName)
	}

	return projectNames
}
