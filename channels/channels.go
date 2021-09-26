package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, signature string) {
	sum := 0
	fmt.Println(s)
	for _, v := range s {
		sum += v
	}
	if signature == "PENDING" {
		time.Sleep(3*time.Second)
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s[:len(s)/2])
	fmt.Println(s[len(s)/2:])
	go sum(s[:len(s)/2], c, "PENDING")
	x := <- c
	go sum(s[len(s)/2:], c, "")
	//fmt.Println(<-c)
	y := <-c // receive from c
	//fmt.Println(<-c)
	fmt.Println(x, y, x+y)
}
