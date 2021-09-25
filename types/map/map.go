package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	//Mutating Map
	testMap := make(map[string]int)

	testMap["Answer"] = 42
	fmt.Println(testMap["Answer"])
	testMap["Answer"] = 18
	fmt.Println(testMap["Answer"])

	delete(testMap, "Answer")
	fmt.Println(testMap["Answer"])

	v, ok := testMap["Answer"]
	fmt.Println(v, ok)
}
