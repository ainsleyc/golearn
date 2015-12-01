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

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return 
}

var outVar bool

func main() {
  var inVar int
  var initVar = "blah" 
  implA, implB := 2, false
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Square root of 4 is", math.Sqrt(4))
	fmt.Println("Pi is", math.Pi)
  fmt.Println("Add", add(1, 2))
  a, b := swap("hello", "world")
  fmt.Println("Swap", a, b) 
  fmt.Println(split(10)) 
  fmt.Println(outVar, inVar, initVar, implA, implB)
}
