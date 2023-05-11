package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefix string
	wordsLetters := make([][]string, len(strs))
	for i, word := range strs {
		wordsLetters[i] = strings.Split(word, "")
	}

	for i, letter := range wordsLetters[0] {
		for _, word := range wordsLetters {
			if i >= len(word) || word[i] != letter {
				return string(prefix)
			}
		}
		prefix += letter
	}

	return prefix
}
