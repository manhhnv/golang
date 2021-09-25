package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	//Like reference to arrays
	names := [4] string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[:2]
	b := names[1: 3]
	fmt.Println(a, b)
	b[0] = "Manh"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%T\n", q)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)

	//Slice length and capacity
	iSlice := []int {2, 3, 5, 7, 11, 13}
	printSlice(iSlice)
	iSlice = iSlice[:0]
	printSlice(iSlice)
	iSlice = iSlice[:4]
	printSlice(iSlice)
	iSlice = iSlice[2:3]
	printSlice(iSlice)

	//Nil slice
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	//Create slice with make function
	aSlice := make([]int, 5)
	printSlice(aSlice)

	bSlice := make([]int, 0, 5)
	printSlice(bSlice)

	cSlice := bSlice[:2]
	printSlice(cSlice)

	dSlice := cSlice[2:5]
	printSlice(dSlice)

	//Slices of slices
	boards := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	boards[0][0] = "X"
	boards[2][2] = "O"
	boards[1][2] = "X"
	boards[1][0] = "O"
	boards[0][2] = "X"
	for i := 0; i < len(boards); i++ {
		fmt.Printf("%s\n", strings.Join(boards[i], " "))
	}
	appendToSlice([]int{})
}

func appendToSlice(s []int)  {
	fmt.Println("Append to slice")
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	a := []int{1, 2, 3, 4, 1, 2, 3, 4}
	s = append(s, a...)
	printSlice(s)

	for i, v := range s {
		fmt.Println(i, v)
	}
}
