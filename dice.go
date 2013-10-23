package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix()) // Seed with time
  fmt.Println(rand.Int()%7) // Print modulo 7 (0 thru 6)
}
