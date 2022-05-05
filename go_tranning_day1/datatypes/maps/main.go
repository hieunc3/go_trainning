package main

import "fmt"

type Employee struct {
	salary  int
	country string
}

func main() {
	//1> create map
	employeeSalary := make(map[string]int)
	employeePosition := map[int]string{
		1: "floor 1",
		2: "floor 2",
		3: "floor 3",
	}
	fmt.Println(employeeSalary)
	fmt.Println(employeePosition)

	//2> add item to a map
	employeeSalary["steve"] = 1200
	employeeSalary["mike"] = 1100
	employeeSalary["lucy"] = 1300
	employeeSalary["jack"] = 900
	employeeSalary["julie"] = 4300
	fmt.Println("EmployeeSalary map contents:", employeeSalary)

	//3> zero value of map
	//run-time panic will occur -> because the map has to be initialize
	//before adding elements.
	// var employeeBounus map[string]int
	// employeeBounus["test"] = 1
	// fmt.Println(employeeBounus)

	//4> retrieve value for a key from a map
	//no runtime error when we try to retrieve the value for a key
	//that is not present in the map
	value, ok := employeeSalary["steven"]
	if ok {
		fmt.Println("Salary of steve:", value)
	} else {
		fmt.Println("Not found")
	}

	//5> iterate over all elements in a map
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	//6> deleting items from a map
	//there will be no runtime error if the delete a key that is
	//not present in the map
	delete(employeeSalary, "steve")

	//7> map of struct
	emp1 := Employee{salary: 1200, country: "USA"}
	emp2 := Employee{salary: 1300, country: "JAPAN"}
	emp3 := Employee{salary: 1400, country: "VIETNAM"}
	listEmployees := map[string]Employee{
		"Tesli":   emp1,
		"William": emp2,
		"Juki":    emp3,
	}
	fmt.Println(listEmployees)

	//8> maps are reference types
	fmt.Println("Original employee salary", employeeSalary)
	modifiedEmployeeSalary := employeeSalary
	modifiedEmployeeSalary["hieu"] = 1235
	modifiedEmployeeSalary["mike"] = 1700
	fmt.Println("Modified employee salary", modifiedEmployeeSalary)
	fmt.Println("Employee salary changed", employeeSalary)
}
