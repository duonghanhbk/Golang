package main

import "fmt"

func modifying() {
	darr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	dSlice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dSlice {
		dSlice[i] += 10
	}
	fmt.Println("array after", darr)
}

func modifying2() {
	numa := [3]int{1, 2, 3}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

func main() {
	modifying()
	modifying2()
}
