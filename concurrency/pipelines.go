package main

import "fmt"

// Defyning read & write channels
func Generator(c chan<- int) { // c chan <- read channel
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(cIn <-chan int, cOut chan<- int) { //cIn channel read chnnael - cOut write chnnael
	for v := range cIn {
		cOut <- v * 2
		// cIn <- 1
	}
	close(cOut)
}

func PrintValues(c <-chan int) {
	for value := range c {
		fmt.Printf("#%d\n", value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)
	go Generator(generator)
	go Double(generator, doubles)
	PrintValues(doubles)
}
