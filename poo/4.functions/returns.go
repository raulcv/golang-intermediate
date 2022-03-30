package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

//Function Variadica
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}
func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	printNames("hola", "hello", "world")
	fmt.Println(getValues(2))
}
