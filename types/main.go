package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	//Pointer
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p /37
	fmt.Println(j)

	//Struct
	fmt.Println(Vertex{X: 20, Y: 10})
	//Struct fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	//Pointers to structs
	q := &v
	q.X = 1e9
	fmt.Println(v)
	//Array
	var a[2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slice
	var s[]int = primes[1:4]
	fmt.Println(s)
}
