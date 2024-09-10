package variablename

import (
	"strings"
	"unicode"
)

func solution(name string) bool {
	for _, namePart := range strings.Split(name, "_") {
		for _, c := range namePart {
			if string(c) == "" {
				continue
			}
			if isEnglishLetter(c) {
				if isDigit(rune(namePart[0])) {
					return false
				}
				continue
			} else if isDigit(c) {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

func isEnglishLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
