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

type TestStruct struct {
  X int
  Y int
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
  ptrVal := 20 
  var ptr = &ptrVal
  fmt.Println("Pounter", *ptr)
  testStruct := TestStruct{1, 2}
  fmt.Println(TestStruct{1, 2})
  testStruct.X = 4
  fmt.Println(testStruct)
  structPtr := &testStruct
  structPtr.X = 3
  fmt.Println(*structPtr)
  var arr[2]int
  arr[0] = 1
  arr[1] = 2
  fmt.Println(arr[0], arr[1])
  arr2 := []int{2, 3, 4}  
  for arrCtr := 0; arrCtr < len(arr2); arrCtr++ {
    fmt.Println(arr2[arrCtr])
  }
  slice := arr2[1:3]
  fmt.Println("Slice", slice[0], slice[1])
  slice2 := append(slice, 5, 6, 7)
  fmt.Println("Slice2", slice2[3], slice2[4])
  for i, v := range slice2 {
    fmt.Printf("Slice2 %d = %d\n", i, v)
  }
  var testMap = map[string]int{
    "ONE": 1,
    "TWO": 2,
  }
  for i, v := range testMap {
    fmt.Printf("Map %s = %d\n", i, v)
  }
}
