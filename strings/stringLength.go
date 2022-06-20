package main

import (
	"fmt"
	"unicode/utf8"
)

func stringLength(s string) {
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(s))
	fmt.Printf("Number of bytes: %d\n", len(s))
}
func main() {
	word1 := "Señor"
	stringLength(word1)
	fmt.Printf("\n")
	word2 := "Pets"
	stringLength(word2)
}
