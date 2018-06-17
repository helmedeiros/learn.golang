package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("nil!")
	}
}
