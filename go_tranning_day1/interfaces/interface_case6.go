package main

import (
	"fmt"
	"runtime"
)

type SalaryCalculator1 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator1
	LeaveCalculator
}

type Employee struct {
	FirstName   string
	LastName    string
	BasicPay    int
	Pf          int
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.FirstName, e.LastName, (e.BasicPay + e.Pf))
}
func (e Employee) CalculateLeavesLeft() int {
	return e.TotalLeaves - e.LeavesTaken
}

func exampleInterface6() {
	//Implementing multiple interfaces
	fmt.Println("The number of logical CPUs usable by current process: ", runtime.NumCPU())

	empl1 := Employee{
		FirstName:   "Hieu",
		LastName:    "Ngueyn",
		BasicPay:    3000,
		Pf:          10,
		TotalLeaves: 15,
		LeavesTaken: 3,
	}
	var s SalaryCalculator1 = empl1
	s.DisplaySalary()
	var l LeaveCalculator = empl1
	fmt.Println("Leaving taken ", l.CalculateLeavesLeft())

	//This is possible since empl1 which of type Employee
	//implement both SalaryCalculator1 & LeaveCalculator interfaces
	var empOp EmployeeOperations = empl1
	empOp.DisplaySalary()
}
