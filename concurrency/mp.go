package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := time.Second * 4
	d2 := time.Second * 2
	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)
	for i := 0; i < 2; i++ { //Multiplexing
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)

		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)
}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(1)
	c <- param
}
