package main

import "fmt"

func main() {
  // Normal declaration
  arrayOfNumbers := [5]int{1, 2, 3, 4, 5}
  // Some kind of lazy declaration...?
  lazilyDeclaredArray := [...]int{1, 3, 5}
  // And apparently key-value pairs, which idk.
  kvpArray := [3]string{1: "Monkey", 2: "Kong"}

  // We can slice like in Python
  slicedNumbers := arrayOfNumbers[0:2]
  fmt.Printf("A slice:\n%v\n", slicedNumbers)

  // Slices are dynamically sized, which is weird
  sliceOfNumbers := arrayOfNumbers[:]
  sliceOfNumbers = append(sliceOfNumbers, 6, 7)
  fmt.Printf("A dynamically allocated slice:\n%v\n", sliceOfNumbers)

  // How you do enumerate():
  println("Printing my lazily declared array:")
  for index, value := range lazilyDeclaredArray {
    fmt.Printf("%d: %d\n", index, value)
  }
  // If you want to skip the use of index, you need to _ the first arg
  // just saying for i := range will net you indices only, not values
  // ---
  // idk what this is about to do.
  fmt.Printf("%s ??? \n", kvpArray[1])
  // sure enough it prints the value at the key 1 lol.
  // You can make other kinds of maps like this:
  //    (build it)   (key type)  (val type)
  //       v          v       v
  myMap := make(map[string]string)
  myMap["dugford"] = "bailey"
  fmt.Printf("%s\n", myMap["dugford"])
}

