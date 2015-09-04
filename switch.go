package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2
  fmt.Print("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("Its the weekend!")
  default:
    fmt.Println("Its a weekday :(")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("its before noon")
  default:
    fmt.Println("its after noon")
  }
}
