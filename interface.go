package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	var b Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	b = &v

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
