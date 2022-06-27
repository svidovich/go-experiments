package main

import (
  "encoding/json"
  "net/http"
  "github.com/google/uuid"
  "time"
)

type DefaultResponse struct {
  Id string `json:"id"`
  Time int64 `json:"time"`
}

func DefaultHandler(writer http.ResponseWriter, request *http.Request) {
  response := DefaultResponse{uuid.New().String(), time.Now().Unix()}
  json.NewEncoder(writer).Encode(response)
}

func main() {
  http.HandleFunc("/default", DefaultHandler)
  http.ListenAndServe(":7070", nil)
}
