package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type orgResponse struct {
	Success      bool           `json:"success"`
	Organization []Organization `json:"data"`
	Message      string         `json:"message"`
}

type Organization struct {
	Name           string      `json:"name"`
	UserName       string      `json:"user_name"`
	Pending        bool        `json:"pending"`
	CreatedDate    createdDate `json:"created_date"`
	Owner          bool        `json:"owner"`
	Uid            string      `json:"uid"`
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	UserEmail      string      `json:"user_email"`
}

type createdDate struct {
	Seconds     int64 `json:"_seconds"`
	NanoSeconds int64 `json:"nano_seconds"`
}

func GetOrganizationList() ([]Organization, error) {
	client := &http.Client{}
	url := "https://auth-module.dev.x.platformer.com/api/v1/organization/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request : %w", err)
	}

	token, err := GetLocallyStoredToken()
	if err != nil {
		return nil, fmt.Errorf("error getting token %s", err)
	}
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

	var orgList orgResponse
	err = json.Unmarshal([]byte(responseString), &orgList)
	if err != nil {
		fmt.Println(err)
	}

	return orgList.Organization, nil
}

func GetOrganizationsNames(organizationList []Organization) []string {
	var organizationsNames []string

	for _, organization := range organizationList {
		organizationsNames = append(organizationsNames, organization.Name)
	}

	return organizationsNames
}
