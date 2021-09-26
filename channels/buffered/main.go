package main

import (
	"fmt"
	"time"
)

func send(c chan int)  {
	for i:= 0; i < 5; i++ {
		fmt.Printf("Send %d to channel \n", i)
		c <- i
	}
}

func receive(c chan int)  {
	for i := 0; i < 5; i++ {
		s := <- c
		fmt.Printf("Receive %d\n", s)
	}
}

func main() {
	c := make(chan int, 3)
	go send(c)
	go receive(c)
	time.Sleep(3 * time.Second)
	fmt.Println("End")
}
