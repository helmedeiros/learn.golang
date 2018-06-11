package main

import (
	"fmt"
	"math"
)

func newt(w float64, x float64) float64 {
	return w - (w*w-x)/(2*w)
}

func sqrt(x float64) (s float64) {
	s = 1.0

	for w := newt(s, x); math.Abs(w-s) > 0.0000000001; {
		s = w
		w = newt(w, x)
	}
	return
}

func main() {
	fmt.Println(sqrt(25))
}
