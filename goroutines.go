package main

import "fmt"

func f(from string) {
  for i := 0; i <= len(from); i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {

  go f("goroutine")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
