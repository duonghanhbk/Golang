package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	emp1 := Employee{
		name:     "James",
		salary:   3000,
		currency: "$",
	}
	emp1.displaySalary()
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())

	p := &c
	fmt.Printf("Area of circle %f", p.Area())

}
