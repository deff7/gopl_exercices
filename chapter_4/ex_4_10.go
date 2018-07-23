package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/deff7/gopl_exercises/chapter_4/github"
)

type filterDates struct {
	Less bool
	Date time.Time
}

var filterDatesFlag filterDates

func (fd *filterDates) Set(val string) error {
	var (
		month = time.Hour * 24 * 31
		year  = month * 12
	)
	switch val {
	case "last_month":
		fd.Less = true
		fd.Date = time.Now().Add(-month)
	case "last_year":
		fd.Less = true
		fd.Date = time.Now().Add(-year)
	case "gt_year_ago":
		fd.Date = time.Now().Add(-year)
	default:
		return fmt.Errorf("unknown value: %s", val)
	}
	return nil
}

func (fd *filterDates) String() string {
	cmp := "greater"
	if fd.Less {
		cmp = "less"
	}
	return fmt.Sprintf("%s %v", cmp, fd.Date)
}

func main() {
	flag.Var(&filterDatesFlag, "date", "(last_month|last_year|gt_year_ago)")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("wrong number of arguments")
	}

	resp, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}

	for _, iss := range resp.Items {
		if !filterDatesFlag.Date.IsZero() {
			cmpFunc := iss.CreatedAt.After
			if filterDatesFlag.Less {
				cmpFunc = iss.CreatedAt.Before
			}
			if cmpFunc(filterDatesFlag.Date) {
				continue
			}
		}
		fmt.Printf("#%-5d %9.9s %50.50s %s\n", iss.Number, iss.User.Login, iss.Title, iss.CreatedAt)
	}
}
