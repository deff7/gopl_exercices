package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/deff7/gopl_exercises/chapter_4/github"
)

var templ *template.Template

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleRepo(w http.ResponseWriter, req *http.Request) {
	args := strings.Split(req.URL.Path, "/")
	if len(args) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if templ == nil {
		err := initTemplate()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
	}

	repo := github.NewRepo(args[2], args[3])
	errors := []error{}

	issues, err := repo.Issues()
	errors = append(errors, err)

	milestones, err := repo.Milestones()
	errors = append(errors, err)

	contributors, err := repo.Contributors()
	errors = append(errors, err)

	ok := true
	for _, err := range errors {
		if err != nil {
			if ok {
				w.WriteHeader(http.StatusInternalServerError)
			}
			ok = false
			fmt.Fprintln(w, err)
		}
	}
	if !ok {
		return
	}

	templ.Execute(w, struct {
		Title        string
		Issues       []github.Issue
		Milestones   []github.Milestone
		Contributors []github.Contributor
	}{
		args[2] + "/" + args[3],
		issues,
		milestones,
		contributors,
	})
}

func initTemplate() error {
	var err error

	f, err := os.Open("ex_4_14.html")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	templ, err = template.New("repo").Parse(string(data))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	http.HandleFunc("/repo/", handleRepo)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
