package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchJSON(u string, dest interface{}) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return fmt.Errorf("%v\n%s", err, string(data))
	}

	return nil
}

type Repo struct {
	Owner, Name string
	url         string
}

func NewRepo(owner, name string) *Repo {
	return &Repo{
		owner,
		name,
		fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name),
	}
}

func (r *Repo) Issues() ([]Issue, error) {
	res := []Issue{}
	err := fetchJSON(r.url+"/issues", &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) Milestones() ([]Milestone, error) {
	res := []Milestone{}
	err := fetchJSON(r.url+"/milestones", &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) Contributors() ([]Contributor, error) {
	res := []Contributor{}
	err := fetchJSON(r.url+"/contributors", &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
