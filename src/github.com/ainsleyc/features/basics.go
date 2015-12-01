package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(x, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
  return y, x 
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Square root of 4 is", math.Sqrt(4))
	fmt.Println("Pi is", math.Pi)
  fmt.Println("Add", add(1, 2))
  a, b := swap("hello", "world")
  fmt.Println("Swap", a, b) 
}
