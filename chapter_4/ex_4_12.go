package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type comics struct {
	ID         int `json:"num"`
	Title      string
	Transcript string
	Alt        string
	ImgURL     string `json:"img"`
}

func fetch(id int) (*comics, error) {
	u := "https://xkcd.com/"
	if id > 0 {
		u += strconv.Itoa(id) + "/"
	}
	u += "info.0.json"

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &comics{}
	if err = json.Unmarshal(data, r); err != nil {
		return nil, err
	}

	return r, nil
}

var (
	indexDir   = "xkcd"
	fileFormat = "json"
)

func formQueue(count int) ([]int, []int, error) {
	dir, err := os.Open(indexDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, nil, err
	}

	if os.IsNotExist(err) {
		r := make([]int, count)
		for i := range r {
			if i+1 == 404 {
				continue
			}
			r[i] = i + 1
		}
		return r, nil, nil
	}

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, nil, err
	}

	present := map[int]bool{}
	for _, file := range files {
		if !strings.HasSuffix(file, fileFormat) {
			continue
		}
		id, err := strconv.Atoi(file[:len(file)-len(fileFormat)-1])
		if err != nil || id > count {
			continue
		}
		present[id] = true
	}

	r, local := []int{}, []int{}
	for id := 1; id <= count; id++ {
		if id == 404 {
			continue
		}
		if present[id] {
			local = append(local, id)
			continue
		}
		r = append(r, id)
	}

	return r, local, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fname(id int) string {
	return fmt.Sprintf("%s/%d.%s", indexDir, id, fileFormat)
}

func buildIndex() map[int]*comics {
	first, err := fetch(0)
	checkError(err)

	queue, localQueue, err := formQueue(first.ID)
	checkError(err)
	if len(queue) > 0 {
		log.Printf("need to fetch %d comics", len(queue))
	}

	index := map[int]*comics{}
	index[first.ID] = first

	err = os.MkdirAll(indexDir, os.ModePerm)
	checkError(err)

	for _, id := range queue {
		c, err := fetch(id)
		checkError(err)

		index[c.ID] = c
		data, err := json.Marshal(c)
		checkError(err)

		err = ioutil.WriteFile(fname(id), data, os.ModePerm)
		checkError(err)

		log.Printf("comics with id = %d saved to local index", id)
	}

	for _, id := range localQueue {
		f, err := os.Open(fname(id))
		checkError(err)

		c := &comics{}
		err = json.NewDecoder(f).Decode(c)
		checkError(err)

		index[c.ID] = c
	}

	return index
}

func main() {
	index := buildIndex()
	terms := os.Args[1:]

	matched := []*comics{}
	for _, c := range index {
		found := false
		for _, t := range terms {
			if found {
				break
			}
			for _, ok := range []bool{
				strings.Contains(c.Title, t),
				strings.Contains(c.Alt, t),
				strings.Contains(c.Transcript, t),
			} {
				if ok {
					found = true
					break
				}
			}

		}
		if found {
			matched = append(matched, c)
		}
	}

	for _, c := range matched {
		fmt.Printf("%s - %50.50s\n", c.ImgURL, c.Transcript)
	}
}
