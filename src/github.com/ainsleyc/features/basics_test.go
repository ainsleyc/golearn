package main

import "testing"

func TestTestStruct(t *testing.T) {
  testStruct := TestStruct{2, 3}
  if testStruct.Sum() != 5 {
    t.Errorf("Incorrect sum: %d + %d = %d", testStruct.X, testStruct.Y, testStruct.Sum())
  }
}
