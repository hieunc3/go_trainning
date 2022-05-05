package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
type Address struct {
	Street   string
	Province string
}
type Student struct {
	Name   string
	Grades []int
	Age    int
	Address
}

func main() {
	fmt.Println("=============== STRUCTS GO LANG ===============")

	hieunguyen := User{Name: "Hieu", Email: "nguyencaohieu26@gmail.com", Status: true, Age: 20}
	//fmt.Printf("Person details are: %+v\n", hieunguyen)
	hieunguyen.printUserEmail()
	hieunguyen.setNewEmail()

	s1 := Student{Name: "Tim", Grades: []int{8, 8, 6}, Age: 20}
	fmt.Println("Student age:", s1.getStudentAge())
	fmt.Println("Average grades of the student:", s1.getAverageGrade())

	//Methods of anonymous struct fields
	s2 := Student{Name: "Tim", Grades: []int{8, 8, 6}, Age: 20, Address: Address{"Tran Phu", "Ha Dong"}}
	//accessing getAddress method of address struct
	s2.getAddress()

	//method accept both pointer & value receiver
	s3 := &s2 //s3 is pointer receiver while s2 is value receiver
	s3.getAddress()

}

//Method in structs
func (u User) printUserEmail() {
	fmt.Println("Your email is: ", u.Email)
}

//Pass the copy of the object into the method -> only change the object's copy email
// func (u User) setNewEmail() {
// 	u.Email = "hieu@gmail.com"
// 	fmt.Println("New email: ", u.Email)
// }

//Pass the reference of the object into method -> change successfully email
func (u *User) setNewEmail() {
	u.Email = "hieu@gmail.com"
	fmt.Println("New email: ", u.Email)
}

func (s Student) getStudentAge() int {
	return s.Age
}

func (s Student) getAverageGrade() float32 {
	var sum float32
	for _, v := range s.Grades {
		sum += float32(v)
	}
	return sum / float32(len(s.Grades))
}

func (a Address) getAddress() {
	fmt.Println("Street", a.Street+" "+a.Province)
}
