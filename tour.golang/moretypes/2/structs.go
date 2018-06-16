package main

import "fmt"

// Vertex is a point in space
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
