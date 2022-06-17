package main

import (
	"fmt"
	"math/rand"
)

func useNormal() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}
}

func useMultipleExpression() {
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func omitExpression() {
	num := 200
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 50 and less than 100", num)
	default:
		fmt.Printf("%d is greater than 100", num)
	}
	fmt.Println()
}

func number() int {
	num := 15 * 5
	return num
}

func fallThrough() {
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

func breakOuter() {
randLoop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randLoop

		default:
			fmt.Println("Generated odd number", i)
		}
	}
}

func main() {
	//useNormal()
	//useMultipleExpression()
	//omitExpression()
	//fallThrough()
	breakOuter()
}
