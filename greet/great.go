package greet 

import ( 
  "fmt"
  "errors"
)  

// Greet or produce error 
func Hello(name string) (string, error) {
  // empty name errors
  if name == "" {
    return "", errors.New("empty name")
  }

  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}
