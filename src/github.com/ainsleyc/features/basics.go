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
  defer fmt.Println("DEFER 1")
  defer fmt.Println("DEFER 2")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Square root of 4 is", math.Sqrt(4))
	fmt.Println("Pi is", math.Pi)
  fmt.Println("Add", add(1, 2))
  a, b := swap("hello", "world")
  fmt.Println("Swap", a, b) 
  fmt.Println(split(10)) 
  fmt.Println(outVar, inVar, initVar, implA, implB)
  i := 3.142
  fmt.Println("Inference", i)
  const Truth = true
  fmt.Println("Const", Truth)
  for i := 1; i < 3; i++ {
    fmt.Println("Loop", i)
  }
  if ifTest := true; ifTest == true {
    fmt.Println("IF")
  }
  
}
