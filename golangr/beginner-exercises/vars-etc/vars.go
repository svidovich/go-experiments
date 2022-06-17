package main
import (
  "fmt"
)

func main() {
  inline_string := "Donkey"
  inline_integer := 5

  var declared_integer int
  var declared_boolean bool

  fmt.Println(declared_boolean)
  fmt.Println(declared_integer)

  fmt.Println(inline_integer)
  fmt.Println(inline_string)

  var is_true = true
  var is_false = false

  is_true, is_false = is_false, is_true

  fmt.Println("This will be false:")

  fmt.Println(is_true)
}
