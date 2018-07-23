package main

func uniqAdjacentStrings(in []string) []string {
	var (
		prev string
		res  = []string{}
	)

	for _, s := range in {
		if prev != "" && prev == s {
			continue
		}
		prev = s
		res = append(res, s)
	}

	return res
}
