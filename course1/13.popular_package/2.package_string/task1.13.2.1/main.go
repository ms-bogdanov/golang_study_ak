package main

import "strings"

func CountWordsInText(txt string, words []string) map[string]int {
	wordTxt := strings.Fields(txt)
	res := map[string]int{}

	for _, word := range words {
		res[word] = 0
	}

	for _, v := range wordTxt {
		if _, ok := res[strings.ToLower(v)]; ok {
			res[strings.ToLower(v)]++
		}
	}
	return res
}
