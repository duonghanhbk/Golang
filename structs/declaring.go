package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person struct {
	string
	int
}

func declaringWithName() {
	emp1 := Employee{
		firstName: "Sam",
		lastName:  "John",
		age:       30,
		salary:    2000,
	}
	emp2 := Employee{
		firstName: "Thomas",
		lastName:  "Jame",
		age:       22,
		salary:    500,
	}
	emp3 := Employee{
		firstName: "Mike",
		lastName:  "Math",
		age:       35,
		salary:    4000,
	}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println("Employee 2", emp3)
}

func declaringAnonymous() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee anonymous", emp3)
}

func anonymousField() {
	p1 := Person{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

func main() {
	var emp4 Employee
	fmt.Println(emp4)
	declaringWithName()
	declaringAnonymous()
	anonymousField()
}
