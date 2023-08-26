package utils

import "strings"

func GetTitle(str string) string {
	wordList := strings.Fields(str)

	for i, word := range wordList {
		wordList[i] = strings.ToTitle(word[:1] + word[1:])
	}

	return strings.Join(wordList, " ")
}
