package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z, w := 1.0, 1.0
	for i := 0; i < 9; i++ {
		if w -= (z*z - x) / (2 * z); w != z {
			z = w
		}
	}
	return z
}

func main() {
	fmt.Println(sqrt(9))
}
