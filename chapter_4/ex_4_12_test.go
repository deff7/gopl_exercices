package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFormQueue(t *testing.T) {
	var (
		err   error
		count = 5
	)

	indexDir, err = ioutil.TempDir("./", "xkcd")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(indexDir)

	local := []int{1, 3, 5}
	for _, i := range local {
		ioutil.WriteFile(fmt.Sprintf("%s/%d.%s", indexDir, i, fileFormat), []byte{}, os.ModePerm)
	}

	r, l, err := formQueue(count)
	if err != nil {
		t.Fatal(err)
	}

	want := []int{2, 4}
	if !reflect.DeepEqual(r, want) {
		t.Errorf("expected formQueue = %v, got %v", want, r)
	}
	if !reflect.DeepEqual(l, local) {
		t.Errorf("expected local queue to eq %v, got %v", local, l)
	}
}
