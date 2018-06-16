package main

import (
	"fmt"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
