package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(Add(20, 30))
	fmt.Println(swap("Singh", "Ashish"))
}

func add(x int, y int) int {
	return x + y
}

func Add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
