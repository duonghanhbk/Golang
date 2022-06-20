package main

import "fmt"

type address struct {
	city  string
	state string
}

type person struct {
	firstName string
	lastName  string
	address
}

type Employee1 struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee1) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee1) changeAge(newAge int) {
	e.age = newAge
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

func main() {
	e := Employee1{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee1 name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee1 name after change: %s", e.name)

	fmt.Printf("\n\nEmployee1 age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee1 age after change: %d\n", e.age)

	p := person{
		firstName: "Hanh",
		lastName:  "Nguyen",
		address: address{
			city:  "HCM",
			state: "VN",
		},
	}

	p.fullAddress()
}
