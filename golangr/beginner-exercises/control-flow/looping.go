package main

import "fmt"
import "math"

func power(x float64, exponent float64) float64 {
  return math.Pow(x, exponent)
}

func main() {
  for i := 0; i < 10; i ++ {
    fmt.Printf("Iteration: %d\n", i)
  }
  fmt.Println(power(2, 8))
}
