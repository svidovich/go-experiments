package main

import (
  "fmt"
  "os"
)

func main() {
  // I think error is a keyword; it seems to highlight in vim
  // I also wish this tutorial would say what the first return
  // value was... ????
  // ALSO, this control flow sux lmao
  if something, err := os.Stat("the_file"); os.IsNotExist(err) { // the fuck is this
    fmt.Printf("The file does not exist! Also, %s\n", something)
  } else {
    fmt.Printf("The file exists! Also %s\n", something)
  }
}
