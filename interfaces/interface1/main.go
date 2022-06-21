package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
func main() {
	pEmp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pEmp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	cEmp1 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	fEmp1 := Freelancer{
		empId:       4,
		ratePerHour: 20,
		totalHours:  20,
	}
	employees := []SalaryCalculator{pEmp1, pEmp2, cEmp1, fEmp1}
	totalExpense(employees)
}
