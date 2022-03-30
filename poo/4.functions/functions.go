package main

import (
	"fmt"
	"time"
)

func mainF() {
	x := 5
	y := func() int { return x * 2 }()
	fmt.Println(y)
	z := 10
	w := func() int { return z * 2 }()
	fmt.Println(w)
	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(time.Second * 5)
		fmt.Println("End function")
		c <- 1
	}()
	<-c
}
