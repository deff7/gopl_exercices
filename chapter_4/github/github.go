// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const APIHost = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCout int `json:"total_count"`
	Items     []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // In markdown format
}

type Milestone struct {
	ID          int
	Title       string
	Description string
}

type Contributor struct {
	ID            int
	Login         string
	Contributions int
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(APIHost + "/search/issues" + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

var token string

func SetToken(t string) {
	token = t
}

func issueRequest(u, verb string, reqBody IssueRequest) error {
	var (
		buf bytes.Buffer
	)
	fmt.Println(u)
	if err := json.NewEncoder(&buf).Encode(&reqBody); err != nil {
		return err
	}

	req, err := http.NewRequest(verb, u, &buf)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

func genURL(owner, repo, id string) string {
	u := fmt.Sprintf("%s/repos/%s/%s/issues", APIHost, owner, repo)
	if id != "" {
		u += "/" + id
	}
	u += "?access_token=" + token
	return u
}

func CreateIssue(owner, repo string, req IssueRequest) error {
	return issueRequest(genURL(owner, repo, ""), "POST", req)
}

func EditIssue(owner, repo, id string, req IssueRequest) error {
	return issueRequest(genURL(owner, repo, id), "PATCH", req)
}

func DeleteIssue(owner, repo, id string) error {
	req := IssueRequest{State: "closed"}
	return issueRequest(genURL(owner, repo, id), "PATCH", req)
}
