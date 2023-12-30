package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	var st = []struct {
		c int
		b bool
		f float32
	}{
		{2, true, 2.1},
		{3, true, 6.7},
		{5, false, 7.9},
	}

	var a = [5]int{1, 2, 3, 4}

	var ptr *[5]int

	ptr = &a

	fmt.Println(ptr)

	fmt.Println(st)

	var p *[]struct {
		c int
		b bool
		f float32
	}

	p = &st

	fmt.Println(p)

	m = make(map[string]Vertex, 10)
	m["Bell Labs"] = Vertex{
		40.490, -74.234,
	}
	fmt.Println(m["Bell Labs"])

	var pt *map[string]Vertex
	pt = &m
	fmt.Println(pt)

	var mt = make(map[string][]int, 10)
	mt["ashish"] = []int{1, 2, 3, 4}
	fmt.Println(mt)
}
