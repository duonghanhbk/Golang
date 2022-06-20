package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

// Passing a slice to a variadic function
func passSlice(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

// Modifying a slice inside variadic function
func change(s ...string) {
	s[0] = "Go"
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	passSlice(89, nums...)

	// Modifying
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
