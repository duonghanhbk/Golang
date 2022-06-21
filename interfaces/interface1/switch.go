package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

type Describe interface {
	Describe()
}

type Manager struct {
	name string
	age  int
}

func (m Manager) Describe() {
	fmt.Printf("%s is %d years old", m.name, m.age)
}

func findTypeInterface(i interface{}) {
	switch v := i.(type) {
	case Describe:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("String")
	findType(443)
	p := Manager{
		name: "Tony",
		age:  32,
	}
	findTypeInterface(p)
}
