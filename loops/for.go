package main

import "fmt"

func useBreak() {
	fmt.Println("Start Break")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d", i)
	}
	fmt.Println()
	fmt.Println("End Break")
}

func useContinue() {
	fmt.Println("Start Continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println("End Continue")
}

func useMore() {
	fmt.Println("Start more")
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
	fmt.Println("End More")
}

func _useInfinitie() {
	for {
		fmt.Println("Hello World")
	}
}

func main() {
	useBreak()
	useContinue()
	useMore()
}
