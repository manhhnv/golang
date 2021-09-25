package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (e *MyError) String() string {
	return "OK"
}

func run() *MyError {
	p := &MyError{
		time.Now(),
		"it didn't work",
	}
	fmt.Println(p.String())
	return p
}

func main() {
	if err := run(); err != nil {
		var a MyError
		a = *err
		fmt.Println(a.String())
		fmt.Println(err)
	}
}
