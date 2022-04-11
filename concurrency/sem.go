package main

import (
	"fmt"
	"sync"
	"time"
)
// Trafig lights
func main() {
	c := make(chan int, 5) // max channel number

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1 // activate like a semaforo
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() // rest -1
	fmt.Printf("Id %d started\n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Id %d finished\n", i)
	<-c // cleanup channel
}
