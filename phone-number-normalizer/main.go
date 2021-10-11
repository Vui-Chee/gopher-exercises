package main

import (
	"strings"
	"unicode"
)

func normalize(phone string) string {
	var builder strings.Builder

	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}

func main() {
}
