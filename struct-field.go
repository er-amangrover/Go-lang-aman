package main

import "fmt"

type Vertex struct {
	X int
	Y float64
}

func main() {
	v := Vertex{1, 2.1}
	v.X = 4
	v1 := Vertex{2, 1.4}
	fmt.Println(v)
	fmt.Println(v1)
}
