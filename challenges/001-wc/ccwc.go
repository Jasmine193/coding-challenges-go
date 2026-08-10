package wc

import (
	"unicode"
	"unicode/utf8"
)

func CountBytes(data []byte) int {
	return len(data)
}

func CountLines(data []byte) int {
	count := 0
	for _, b := range data {
		if b == '\n' {
			count++
		}
	}
	return count
}

func CountWords(data []byte) int {
	count := 0
	inWord := false

	for _, b := range data {
		if unicode.IsSpace(rune(b)) {
			inWord = false
		}
		if inWord != true && !unicode.IsSpace(rune(b)) {
			count++
			inWord = true
		}
	}
	return count
}

func CountCharacters(data []byte) int {
	return utf8.RuneCount(data)
}
