package main

// Vertex Represents a corner or a point where lines meet
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  //has type Vertex
	v2 = Vertex{X: 1}  //Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} //hat type *Vertex
)
