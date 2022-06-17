package main

import "fmt"

func difference() {
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Println("length of slice capacity", len(fruitSlice), cap(fruitSlice)) //length of fruitSlice is 2 and capacity is 6
}

func reSliced() {
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitSlice), cap(fruitSlice)) //length of is 2 and capacity is 6
	fruitSlice = fruitSlice[:cap(fruitSlice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitSlice), "and capacity is", cap(fruitSlice))
}

func main() {
	difference()
	reSliced()
}
