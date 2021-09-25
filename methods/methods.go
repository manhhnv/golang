package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64  {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(2)
	fmt.Println(v.Abs())

	testVertex := Vertex{5, 6}
	p := &testVertex
	fmt.Println(p.Abs())
	fmt.Println(Abs(*p))

	c := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", c, c.Abs())
	c.Scale(2)
	fmt.Printf("After scaling: %+v, Abs: %v\n", c, c.Abs())
}
