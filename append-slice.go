package main

import "fmt"

func main() {
	var s []int
	PrintSlice(s)
	for i := 0; i < 17; i++ {
		s = append(s, i)
		PrintSlice(s)
	}
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
