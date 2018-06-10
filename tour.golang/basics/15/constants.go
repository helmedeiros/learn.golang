package main

import "fmt"

//Pi is a raw precision for Math.Pi
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Is the truth out there?", Truth)
}
