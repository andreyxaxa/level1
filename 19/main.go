package main

import (
	"fmt"
)

func reverseString(original string) string {
	originalRuned := []rune(original)

	for i := 0; i < len(originalRuned)/2; i++ {
		originalRuned[i], originalRuned[len(originalRuned)-i-1] = originalRuned[len(originalRuned)-i-1], originalRuned[i]
	}

	return string(originalRuned)
}

func main() {
	s := "главрыба"

	fmt.Println(reverseString(s))
}
