package main

import (
	"fmt"
)

func main() {
	if num := 11; num%2 == 0 { //checks if number is even
		fmt.Println("the number", num, "is even")
	} else {
		fmt.Println("the number", num, "is odd")
	}
}
