package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type movie struct {
	Title     string
	PosterURL string `json:"Poster"`
	Error     string
}

func requestURL(title, key string) string {
	return fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", key, title)
}

func fetch(u string) ([]byte, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fetchMovie(title, key string) (*movie, error) {
	data, err := fetch(requestURL(title, key))
	if err != nil {
		return nil, err
	}

	mov := &movie{}
	err = json.Unmarshal(data, mov)
	if err != nil {
		return nil, err
	}

	if mov.Error != "" {
		return nil, fmt.Errorf("api error: %s", mov.Error)
	}

	if mov.PosterURL == "" || mov.PosterURL == "N/A" {
		return nil, fmt.Errorf("no poster")
	}

	return mov, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func snakeCase(src string) string {
	ss := strings.Split(src, " ")
	for i, s := range ss {
		ss[i] = strings.ToLower(s)
	}
	return strings.Join(ss, "_")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./poster title")
		os.Exit(1)
	}

	var (
		title = strings.Join(os.Args[1:], " ")
		key   = os.Getenv("OMDB_KEY")
	)

	if key == "" {
		fmt.Println("OMDB_KEY variable is not set")
		os.Exit(1)
	}

	mov, err := fetchMovie(title, key)
	checkError(err)

	poster, err := fetch(mov.PosterURL)
	checkError(err)

	var (
		fname = snakeCase(mov.Title)
		ext   = filepath.Ext(mov.PosterURL)
	)

	ioutil.WriteFile(fname+"."+ext, poster, os.ModePerm)
}
