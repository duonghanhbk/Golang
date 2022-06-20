package main

import "fmt"

func main() {
	value := 255
	var a *int
	if a == nil {
		fmt.Println("a is", a)
		a = &value
		fmt.Println("a after initialization is", a)
	}

}
