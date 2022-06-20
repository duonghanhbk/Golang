package main

import "fmt"

type Employee1 struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp := &Employee1{
		firstName: "Sam",
		lastName:  "Ho",
		age:       32,
		salary:    200,
	}
	fmt.Println(emp)
	fmt.Println("First Name:", (*emp).firstName, emp.firstName)
	fmt.Println("Age:", (*emp).age, emp.age)
}
