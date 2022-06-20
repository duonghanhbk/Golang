package main

import "fmt"

// create map
//func create() map[string]int {
//	return make(map[string]int)
//}

func singleMap() {
	employeeSalary := map[string]int{
		"jamie": 5000,
		"mike":  3000,
	}
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		employeeSalary[newEmp] = 2000
		fmt.Println(newEmp, "not found")
	}

	for key, value := range employeeSalary {
		fmt.Printf("Salary of [%s] = %d\n", key, value)
	}

	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "mike")
	fmt.Println("map after deletion", employeeSalary)
}

// create map of structs
type employee struct {
	salary  int
	country string
}

func mapOfStructs() {
	emp1 := employee{
		salary:  1000,
		country: "India",
	}
	emp2 := employee{
		salary:  2000,
		country: "USA",
	}
	emp3 := employee{
		salary:  4000,
		country: "VN",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Joe":   emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s, Salary: $%d, Country: %s\n", name, info.salary, info.country)
	}
}

// Modifying
func modifying() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)
}
func main() {
	singleMap()
	mapOfStructs()
	modifying()
}
