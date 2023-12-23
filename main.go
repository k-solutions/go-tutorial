package main

import (
  "errors"
	"fmt"
  "example.com/greet"
)

func main() {
  // empty name errors
  if name == "" {
    return "", errors.New("empty name")
  }
  message := greet.Hello("Test")
	fmt.Println(message)
}
