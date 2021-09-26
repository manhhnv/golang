package main

import (
	"fmt"
	"sync"
	"time"
)

func say(message string, n int, waitGroup *sync.WaitGroup)  {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(message)
	}
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go say("Message 1", 10, &waitGroup)
	go say("Message 2", 2, &waitGroup)
	waitGroup.Wait()
	fmt.Println("Main func is done")
}
