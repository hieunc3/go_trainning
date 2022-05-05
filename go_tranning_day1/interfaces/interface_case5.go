package main

import "fmt"

type Describer interface {
	Describe()
}

type PersonInfo struct {
	Name string
	Age  int
}

type AddressInfo struct {
	State   string
	Country string
}

//implement using value receiver
func (p PersonInfo) Describe() {
	fmt.Printf("%s is %d years old\n", p.Name, p.Age)
}

//implement using pointer receiver
//Describer interface is implemented by AddressInfo pointer
func (a *AddressInfo) Describe() {
	fmt.Printf("State %s & Country %s", a.State, a.Country)
}

func exampleInterface5() {
	var d1 Describer
	per1 := PersonInfo{"Sam", 20}
	d1 = per1
	d1.Describe()

	per2 := PersonInfo{"James", 30}
	d1 = &per2
	d1.Describe()

	var d2 Describer
	a := AddressInfo{"Washington", "USA"}
	a.Describe()
	//Cannot use a type AddressInfo as type Describer
	//Address does not implement Describer
	//Describe method has pointer receiver
	/*
		The reason is that it's legal to call a pointer - value method
		on anything that is alrady a pointer or whose address can be taken.
		The concrete value stored in an interface is not addressable
		and hence it is not possible for the complier to automatically
		taken the address of a in line 53
	*/
	//d2= a
	d2 = &a
	d2.Describe()
}
