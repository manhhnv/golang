package main

import (
	"fmt"
	"time"
)

func say(message string, n int)  {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(message)
	}
}

func main() {
	go say("Message 1", 10)
	say("Message 2", 2)
}
