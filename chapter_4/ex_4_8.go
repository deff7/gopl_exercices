package main

import "unicode"

type category struct {
	Name string
	Func func(rune) bool
}

var categories = []category{
	{"letter", unicode.IsLetter},
	{"digit", unicode.IsDigit},
	{"mark", unicode.IsMark},
	{"other", func(rune) bool { return true }},
}

type countMap map[string]map[rune]int

func charcount(s string) countMap {
	count := countMap{}
	for _, r := range s {
		for _, cat := range categories {
			if cat.Func(r) {
				if count[cat.Name] == nil {
					count[cat.Name] = map[rune]int{}
				}
				count[cat.Name][r]++
				break
			}
		}
	}
	return count
}
