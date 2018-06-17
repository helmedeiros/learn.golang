package main

import (
	"fmt"
)

// Hello will greet someone
func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return "Hello, " + s
}

func main() {
	fmt.Println(Hello("Chris"))
}
