package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func countLines(r io.Reader, file string, counts map[string]int, foundIn map[string]map[string]bool) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		counts[line]++
		if foundIn[line] == nil {
			foundIn[line] = make(map[string]bool)
		}
		foundIn[line][file] = true
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		counts  = make(map[string]int)
		foundIn = make(map[string]map[string]bool)
		files   = os.Args[1:]
	)

	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		countLines(f, fname, counts, foundIn)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file := range foundIn[line] {
				fmt.Printf("\t%s\n", file)
			}
		}
	}
}
