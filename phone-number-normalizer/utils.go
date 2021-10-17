package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func scream(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func normalize(phone string) string {
	var builder strings.Builder
	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
