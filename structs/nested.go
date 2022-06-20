package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Person1 struct {
	name    string
	age     int
	address Address
}

func main() {
	p := Person1{
		name: "Nol",
		age:  33,
		address: Address{
			city:  "HCM",
			state: "VN",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}
