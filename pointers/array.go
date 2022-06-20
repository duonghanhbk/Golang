package main

import "fmt"

func useArrayModify(arr *[3]int) {
	arr[0] = 100
}

func useSliceModify(slc []int) {
	slc[0] = 200
}
func main() {
	a := [3]int{20, 30, 40}
	useArrayModify(&a)
	fmt.Println(a)
	b := []int{11, 22, 33}
	useSliceModify(b[:])
	fmt.Println(b)
}
