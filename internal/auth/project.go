package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

const projectListURL = "https://auth-module.dev.x.platformer.com/api/v1/organization/project/list"

// Project models a Platformer Project
type Project struct {
	ID    string `json:"projectId"`
	Name  string `json:"projectName"`
	Roles []struct {
		CollectionID string   `json:"collection_id"`
		RolesID      []string `json:"roles_id"`
	} `json:"roles"`
	IsProjectOwner bool `json:"isProjectOwner"`
}

// ProjectList is a map of Projects by name
type ProjectList map[string]Project

// Names returns the names of the projects
func (p ProjectList) Names() []string {
	names := []string{}
	for n := range p {
		names = append(names, n)
	}
	return names
}

// LoadProjectList loads a list of Projects for the given organization
func LoadProjectList(organizationID string) (ProjectList, error) {
	url := fmt.Sprintf("%s/%s", projectListURL, organizationID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", config.GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var projectListResponse struct {
		Data []Project `json:"data"`
	}

	if err = json.Unmarshal(body, &projectListResponse); err != nil {
		return nil, err
	}

	projectList := ProjectList{}
	for _, p := range projectListResponse.Data {
		projectList[p.Name] = p
	}

	return projectList, nil
}
