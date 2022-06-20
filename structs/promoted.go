package main

import "fmt"

type AddressP struct {
	city  string
	state string
}
type PersonP struct {
	name string
	age  int
	AddressP
}

func main() {
	p := PersonP{
		name: "Naveen",
		age:  50,
		AddressP: AddressP{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
