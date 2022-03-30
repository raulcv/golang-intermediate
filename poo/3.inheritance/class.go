package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	id int
}
type FullTimeEmployee struct {
	Person
	Employee
}

func (fullTimeEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxtRate int
}

func (temporaryEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(pi PrintInfo) {
	fmt.Println(pi.getMessage())
}
func main() {
	fullTimeEmployee := FullTimeEmployee{}
	fullTimeEmployee.name = "Name"
	fullTimeEmployee.age = 25
	fullTimeEmployee.id = 1
	fmt.Printf("Id: %d, Name: %s, Vacation: %d\n", fullTimeEmployee.id, fullTimeEmployee.name, fullTimeEmployee.age)
	temporaryEmployee := TemporaryEmployee{}
	getMessage(fullTimeEmployee)
	getMessage(temporaryEmployee)
}
