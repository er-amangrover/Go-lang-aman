package main

import "fmt"

func main() {
	a := make([]int, 5)
	PrintSlice(a)

	a = make([]int, 10)
	PrintSlice(a)

	b := make([]int, 0, 5)
	PrintSlice(b)

}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
