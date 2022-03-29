package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id: id, name: name, vacation: vacation}
}
func main() {
	//Forma 1
	e1 := Employee{}
	fmt.Println(e1)
	//Forma 2
	e2 := Employee{id: 1, name: "Raul", vacation: true}
	fmt.Printf("%v\n", e2)
	fmt.Printf("Id: %d\nName: %s\nVacation: %t\n", e2.id, e2.name, e2.vacation)
	//Forma 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 12
	e3.name = "Juan"
	e3.vacation = true
	fmt.Printf("Id: %d, Name: %s, Vacation: %t\n", e3.id, e3.name, e3.vacation)
	//Forma 4
	e4 := NewEmployee(15, "July", true)
	fmt.Printf("Id: %d, Name: %s, Vacation: %t\n", e4.id, e4.name, e4.vacation)

}
