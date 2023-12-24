package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.400, -74.234,
	},
	"Google": Vertex{
		37.422, -122.089,
	},
}

func main() {
	fmt.Println(m)
}
