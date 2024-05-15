package transformers

func FormatPunctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	// Move punctuation to the previous word if it's in the middle of a string
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}
	// Move punctuation to the previous word if it's at the end of the string
	for i, word := range s {
		for _, punc := range puncs {
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}
	// Remove punctuation if it's in the middle of a string and not part of a group
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	// Handle apostrophes
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
