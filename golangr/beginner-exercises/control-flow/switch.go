package main

import (
  "fmt"
  "time"
)

func main() {
  now := time.Now().Unix()
  fmt.Println(now)
  // It's good to have switches back.
  switch {
    case now % 2 == 0:
      fmt.Printf("The current time, %d, is even\n", now)
    case now % 2 != 0:
      fmt.Printf("The current time, %d, is odd\n", now)
    default:
      fmt.Printf("The current time, %d, breaks mathematics\n", now)
  }
}
