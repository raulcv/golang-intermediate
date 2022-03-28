package main

import (
	"fmt"
	"strconv"
	"time"
)

func repasoGoBasicFunction() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("125900", 0, 32)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	m := make(map[string]int)
	m["Key"] = 6
	m["map"] = 7
	fmt.Println(m["Key"]) //Print 6
	s := []string{"Raul", "Bianch", "Jessie"}
	// for index, value := range s {
	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// }
	s = append(s, "Other People")
	for index, value := range s {
		fmt.Printf("%d: %s\n", index, value)
	}
	//Canales para goRoutines
	// c := make(chan string)
	// go doSomething(c)
	// <-c
	//Apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan<- string) {
	time.Sleep(time.Second * 3)
	fmt.Println("Done!")
	c <- "Foreing friend"
}
