package main

func anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == "" {
		return false
	}

	var (
		chars = map[rune]int{}
		ra    = []rune(a)
		rb    = []rune(b)
	)

	for i, c := range ra {
		chars[c]++
		chars[rb[i]]--
	}

	for _, c := range chars {
		if c != 0 {
			return false
		}
	}

	return true
}
