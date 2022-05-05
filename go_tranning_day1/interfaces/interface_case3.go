package main

import "fmt"

type Worker interface {
	Work()
}
type Person struct {
	Name string
}

func (p Person) Work() {
	fmt.Println(p.Name + " is working")
}
func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func exampleInterface3() {
	//we assign w variable type Worker equal to p1
	//now the concrete type of w is Person with name field
	//the func describe prints the value and concrete type of
	//the interface
	p1 := Person{
		Name: "Hieu Nguyen",
	}
	var w Worker = p1
	describe(w)
	w.Work()
}
