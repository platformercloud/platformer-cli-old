package auth

import (
	"encoding/json"
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/util"
	"io/ioutil"
	"net/http"

	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/platformercloud/platformer-cli/internal/config"
)

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
	var names []string
	for n := range p {
		names = append(names, n)
	}
	return names
}

// LoadProjectList loads a list of Projects for the given organization
func LoadProjectList(organizationID string) (ProjectList, error) {
	url := fmt.Sprintf("%s/%s", util.AuthProjectListURL, organizationID)
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

// LoadProjectsFromDefaultOrFlag loads the projects for either the provided
// orgFlag or the default organization. The returned error will be of type cli.Error
// Note: If orgFlag is provided it will not use the saved default organization.
func LoadProjectsFromDefaultOrFlag(orgFlag string) (orgName string, projectList ProjectList, err error) {
	orgList, err := LoadOrganizationList()
	if err != nil {
		return "", nil, &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	if orgFlag != "" {
		// The org flag has been set; ignore the currently selected org in the local config
		if _, ok := orgList[orgFlag]; !ok {
			return "", nil, &cli.UserError{Message: fmt.Sprintf("the organization [%s] does not exist "+
				"or you do not have permission to access it", orgFlag)}
		}
		orgName = orgFlag
	} else if defaultOrgName := config.GetDefaultOrg(); defaultOrgName != "" {
		if _, ok := orgList[defaultOrgName]; !ok {
			return "", nil, &cli.UserError{Message: fmt.Sprintf("the currently selected organization [%s] does not exist "+
				"or you do not have permission to access it."+
				"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
				defaultOrgName),
			}
		}
		orgName = defaultOrgName
	}

	if orgName == "" {
		return "", nil, &cli.UserError{Message: "an organization must be selected." +
			"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
		}
	}

	projectList, err = LoadProjectList(orgList[orgName].ID)
	if err != nil {
		return orgName, nil, &cli.InternalError{Err: err, Message: "failed to load project data"}
	}

	return orgName, projectList, nil
}

// GetProjectIDFromName returns the *Project from a given name and an Organization ID
func GetProjectIDFromName(orgID string, projectName string) (*Project, error) {
	projectList, err := LoadProjectList(orgID)
	if err != nil {
		return nil, err
	}

	for _, n := range projectList.Names() {
		if projectName == n {
			p := projectList[n]
			return &p, nil
		}
	}

	return nil, fmt.Errorf("project does not exist")
}
