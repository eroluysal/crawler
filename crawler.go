package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Status struct {
	CurrentUserUrl                   string `json:"current_user_url"`
	CurrentUserAuthorizationsHtmlUrl string `json:"current_user_authorizations_html_url"`
	AuthorizationsUrl                string `json:"authorizations_url"`
	CodeSearchUrl                    string `json:"code_search_url"`
	CommitSearchUrl                  string `json:"commit_search_url"`
	EmailsUrl                        string `json:"emails_url"`
	EmojisUrl                        string `json:"emojis_url"`
	EventsUrl                        string `json:"events_url"`
	FeedsUrl                         string `json:"feeds_url"`
	FollowersUrl                     string `json:"followers_url"`
	FollowingUrl                     string `json:"following_url"`
	GistsUrl                         string `json:"gists_url"`
	HubUrl                           string `json:"hub_url"`
	IssueSearchUrl                   string `json:"issue_search_url"`
	KeysUrl                          string `json:"keys_url"`
	NotificationsUrl                 string `json:"notifications_url"`
	OrganizationRepositoriesUrl      string `json:"organization_repositories_url"`
	OrganizationUrl                  string `json:"organization_url"`
	PublicGistsUrl                   string `json:"public_gists_url"`
	RateLimitUrl                     string `json:"rate_limit_url"`
	RepositoryUrl                    string `json:"repository_url"`
	RepositorySearchUrl              string `json:"repository_search_url"`
	CurrentUserRepositoriesUrl       string `json:"current_user_repositories_url"`
	StarredUrl                       string `json:"starred_url"`
	StarredGistsUrl                  string `json:"starred_gists_url"`
	TeamUrl                          string `json:"team_url"`
	UserUrl                          string `json:"user_url"`
	UserOrganizationsUrl             string `json:"user_organizations_url"`
	UserRepositoriesUrl              string `json:"user_repositories_url"`
	UserSearchUrl                    string `json:"user_search_url"`
}

const ApiEndpoint string = "https://api.github.com/"

func main() {
	values := Status{}

	response, err := http.Get(ApiEndpoint)
	exitWithMessageIfExistErr(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	exitWithMessageIfExistErr(err)

	json.Unmarshal(body, &values)

	fmt.Printf("Response: \n%s", values)
}

func exitWithMessageIfExistErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
