package transformers

import "unicode"

func ToTitle(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(unicode.ToTitle(rune(s[0]))) + s[1:]
}
