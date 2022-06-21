package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Empty interface {
	ShowEmpty()
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	s := SalaryCalculator(e)
	s.DisplaySalary()
	l := LeaveCalculator(e)
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	// Embadding
	o := EmployeeOperations(e)
	o.DisplaySalary()
	fmt.Println("\nLeaves left =", o.CalculateLeavesLeft())

	var empty Empty
	if empty == nil {
		fmt.Printf("empty is nil and has type %T value %v\n", empty, empty)
	}
}
