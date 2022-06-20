package main

import "fmt"

// Pass pointer to function
func change(val *int) {
	*val = 55
}

// Return pointer from a function
func hello() *int {
	i := 5
	return &i
}

func main() {
	a := 33
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
	c := hello()
	fmt.Println("C value is", *c)
}
