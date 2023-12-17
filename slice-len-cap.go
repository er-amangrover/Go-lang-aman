package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	PrintSlice(s)

	// extend it's length
	s = s[:4]
	PrintSlice(s)

	// drop  its first two value
	s = s[2:]
	PrintSlice(s)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
