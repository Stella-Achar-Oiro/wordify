package transformers

import "strings"

func CorrectArticles(words []string) []string {
	vowels := "aeiouhAEIOUH" // Vowels and 'h' in both cases
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		nextWord := words[i+1]
		if (word == "a" || word == "A") && len(nextWord) > 0 && strings.ContainsRune(vowels, rune(nextWord[0])) {
			if word == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return words
}
