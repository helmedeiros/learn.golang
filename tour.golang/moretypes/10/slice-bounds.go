package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	primes = primes[1:4]
	fmt.Println(primes)

	primes = primes[:2]
	fmt.Println(primes)

	primes = primes[1:]
	fmt.Println(primes)
}
