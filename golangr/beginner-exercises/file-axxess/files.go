package main

import (
  "fmt"
  "os"
  "io/ioutil" // How does this work?
)

func main() {
  file_name := "the_file"
  // I think error is a keyword; it seems to highlight in vim
  // I also wish this tutorial would say what the first return
  // value was... ????
  // ALSO, this control flow sux lmao
  if something, err := os.Stat(file_name); os.IsNotExist(err) { // the fuck is this
    fmt.Printf("The file does not exist! Also, %s\n", something)
  } else {
    file_bytes, err := ioutil.ReadFile(file_name)
    if err != nil {
      fmt.Printf("Error reading file: %s", err)
      return
    }
    file_bytes_as_string := string(file_bytes)
    fmt.Println(file_bytes_as_string)
  }
}
