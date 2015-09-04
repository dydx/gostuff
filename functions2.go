package main

import "fmt"

func vals() (int, int) {
  // how do I idiomatically return error statements?
  // in Node, its expected to pass `err` as the first
  // parameter in the callback... is there anything like
  // this in Go?
  return 3, 7
}

func main() {
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
}
