package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Joke struct {
  // For encoding/json to access the fields of this struct,
  // they need to be capitalized... but the response JSON may
  // not have these field names, so we make a sort of mapping
  // by adding `json:"id"`, for example, to the Id field of
  // the struct.
  Id string `json:"id"`
  Joke string `json:"joke"`
  Status int `json:"status"`
}

func main () {
  jokeAPIURI := "https://icanhazdadjoke.com/"
  client := &http.Client{}
  request, _ := http.NewRequest("GET", jokeAPIURI, nil)
  request.Header.Add("Accept", "application/json")
  response, err := client.Do(request)
  if (err != nil) {
    fmt.Printf("Sorry bud, Dad couldn't find a joke for you.\n")
  } else {
    defer response.Body.Close()
    bodyJSON, err := ioutil.ReadAll(response.Body)
    if (err != nil) {
       fmt.Printf("Couldn't parse the response: %s", err)
    }
    var jokeObject Joke
    decodeError := json.Unmarshal(bodyJSON, &jokeObject)
    if (decodeError != nil) {
       fmt.Printf("Can't decode: %s", decodeError)
    }
    fmt.Printf("%s\n", jokeObject.Joke)
  }
}
