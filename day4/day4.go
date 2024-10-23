package day4

import (
	"slices"
	"strings"
)

func P1(input string) int {
	return validatePassphrases(input, allWordsUnique)
}

func P2(input string) int {
	return validatePassphrases(input, func(words []string) bool {
		sortedCharactersOfWords := make([]string, 0)
		for _, word := range words {
			//take the characters of each word and sort them and then turn them into a string
			runeSlice := make([]rune, 0)
			for _, rune := range word {
				runeSlice = append(runeSlice, rune)
			}
			slices.Sort(runeSlice)
			sortedCharactersOfWords = append(sortedCharactersOfWords, string(runeSlice))
		}
		anagramMap := make(map[string]int)
		for _, sortedCharacterWord := range sortedCharactersOfWords {
			if _, present := anagramMap[sortedCharacterWord]; present {
				return false
			} else {
				anagramMap[sortedCharacterWord] = 1
			}
		}
		return true
	})
}

func validatePassphrases(passphrases string, lineValidator func([]string) bool) int {
	lines := strings.Split(passphrases, "\n")
	validCount := 0
	for _, line := range lines {
		words := strings.Fields(line)
		if lineValidator(words) {
			validCount += 1
		}
	}
	return validCount
}

func allWordsUnique(words []string) bool {
	wordMap := make(map[string]int)
	for _, word := range words {
		if _, present := wordMap[word]; present {
			return false
		} else {
			wordMap[word] = 1
		}
	}
	return true
}
