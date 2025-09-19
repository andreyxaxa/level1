package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(original string) string {
	// слова разделяются одиночным пробелом
	words := strings.Split(original, " ")
	slices.Reverse(words)

	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"

	fmt.Println(reverseWords(s))
}
