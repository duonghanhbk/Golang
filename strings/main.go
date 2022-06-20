package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, value := range s {
		fmt.Printf("%c starts at byte %d\n", value, index)
	}
}

func fromSliceOfBytesHex() {
	bytes := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(bytes)
	fmt.Println(str)
}

func fromSliceBytesDec() {
	bytes := []byte{67, 97, 102, 195, 169}
	str := string(bytes)
	fmt.Println(str)
}

func fromSliceRunes() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	charsAndBytePosition(name)
	fmt.Println()
	fromSliceOfBytesHex()
	fmt.Println()
	fromSliceBytesDec()
	fmt.Println()
	fromSliceRunes()
}
