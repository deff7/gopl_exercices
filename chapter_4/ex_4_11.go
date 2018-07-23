package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/deff7/gopl_exercises/chapter_4/github"
)

const (
	stScanAccessToken = iota
	stScanOwner
	stScanRepo
	stScanAction
	stScanID
	stScanTitle
	stScanBody
	stPerform
	stQuit
)

var prompt = map[int]string{
	stScanAccessToken: "Your access token",
	stScanOwner:       "Owner",
	stScanRepo:        "Repository",
	stScanAction:      "What you want to do? (c - create, e - edit, d - delete)",
	stScanID:          "ID of issue",
	stScanTitle:       "Issue title",
	stScanBody:        "Issue text",
	stPerform:         "Perform action? (y/n)",
}

func createIssue(owner, repo, title, body string) {
	r := github.IssueRequest{Title: title, Body: body}
	if err := github.CreateIssue(owner, repo, r); err != nil {
		log.Fatal(err)
	}
}

func editIssue(owner, repo, id, title, body string) {
	r := github.IssueRequest{Title: title, Body: body}
	if err := github.EditIssue(owner, repo, id, r); err != nil {
		log.Fatal(err)
	}
}

func deleteIssue(owner, repo, id string) {
	if err := github.DeleteIssue(owner, repo, id); err != nil {
		log.Fatal(err)
	}
}

func getAccessToken() string {
	f, err := os.Open(".access_token")
	if err != nil && os.IsNotExist(err) {
		return ""
	}
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func writeAccessToken(token string) {
	f, err := os.Create(".access_token")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(token)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		owner, repo, action string
		id, title, body     string

		token = getAccessToken()
		state = stScanOwner
		err   error
	)

	if len(os.Args) == 3 {
		owner = os.Args[1]
		repo = os.Args[2]
		state = stScanAction
	}

	initState := state
	if token == "" {
		state = stScanAccessToken
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if state == stQuit {
			break
		}

		fmt.Printf("%s: ", prompt[state])

		if !scanner.Scan() {
			err = scanner.Err()
			return
		}
		var (
			s = scanner.Text()
		)

		switch state {
		case stScanAccessToken:
			token = s
			writeAccessToken(token)
			state = initState
		case stScanOwner:
			owner = s
			state = stScanRepo
		case stScanRepo:
			repo = s
			state = stScanAction
		case stScanAction:
			if len(s) > 1 || !strings.Contains("cde", s) {
				break
			}
			action = s
			switch action {
			case "c":
				state = stScanTitle
			case "e", "d":
				state = stScanID
			}
		case stScanID:
			id = s
			if action == "d" {
				state = stPerform
				break
			}
			state = stScanTitle
		case stScanTitle:
			title = s
			state = stScanBody
		case stScanBody:
			body = s
			state = stPerform
		case stPerform:
			if len(s) > 1 || !strings.Contains("yn", s) {
				break
			}
			state = stQuit
			if s == "n" {
				break
			}
			github.SetToken(token)
			switch action {
			case "c":
				createIssue(owner, repo, title, body)
			case "e":
				editIssue(owner, repo, id, title, body)
			case "d":
				deleteIssue(owner, repo, id)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
