package main

import "fmt"

type Rectangle struct {
  length int
  width int
}

func area(r Rectangle) int {
  return r.length * r.width
}

func main() {
  var myRectangle Rectangle
  myRectangle.length = 5
  myRectangle.width = 7
  myRectangleArea := area(myRectangle)
  fmt.Printf("My Rectangle's Area is %d\n", myRectangleArea)
}
