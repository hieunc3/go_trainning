package main

import "fmt"

//declare data type SalaryCalculator interface with method CalculateSalary int
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	EmpId    int
	Basicpay int
	Pf       int
}

type Contract struct {
	EmpId    int
	Basicpay int
}

//salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.Basicpay + p.Pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.Basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
func exampleInterface2() {
	//We have two kinds of employees in the company: Permanent & Contract
	//We have the different between two employees in how to calculate salary
	//By declaring the method, both Permanent & Contract structs
	//now implement the SalaryCalculator interface
	pemp1 := Permanent{
		EmpId:    1,
		Basicpay: 1000,
		Pf:       30,
	}
	pemp2 := Permanent{
		EmpId:    2,
		Basicpay: 1000,
		Pf:       20,
	}
	cemp1 := Contract{
		EmpId:    3,
		Basicpay: 1000,
	}
	//Create a slice of type SalaryCalculator
	//the slice we pass contains both Permanent & Contract types to
	//the totalExpense func. And the func call the CalculateSalary
	//method of the corresponding type
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

/*
	No matter how many type of employee we added. We don't need to
	modify the func totalExpense because the new employee will also
	implement the Salarycalculator interface

*/
