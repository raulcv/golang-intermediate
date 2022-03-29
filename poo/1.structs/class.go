package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}
func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) SetName(name string) {
	e.name = name
}
func (e *Employee) GetName() string {
	return e.name
}
func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "Raul Valeriano"
	e.SetId(5)
	fmt.Printf("%v\n", e)
	e.SetName("PAUL JOHN")
	fmt.Printf("%v\n", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
