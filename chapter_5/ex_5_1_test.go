package main

import (
	"os"
	"reflect"
	"testing"
)

func TestFindlinks(t *testing.T) {
	f, err := os.Open("test_page.html")
	if err != nil {
		t.Fatal(err)
	}
	links, err := findlinks(f)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"foo",
		"bar",
		"yoo",
	}
	if !reflect.DeepEqual(links, want) {
		t.Errorf("got %v", links)
	}
}
