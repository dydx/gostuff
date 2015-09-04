package main

import "fmt"

func main() {
  // map[key-type]val-type
  m := make(map[string]int)

  // set [int] vals to [string] keys in `map`
  m["k1"] = 7
  m["k2"] = 13

  // show the whole map
  fmt.Println("map:", m)

  // get a value back out of the map via [string] key
  v1 := m["k1"]
  fmt.Println("v1: ", v1)

  // get the length/size of the map
  fmt.Println("len:", len(m))

  // remove a key and value from the map
  delete(m, "k2")
  fmt.Println("map:", m)

  // the second return value from retrieving a value
  // indicates whether the item was present or not
  // NOTE: `_` is a blank identifier, for when you don't
  // care about the value of that return item
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  val, prs := m["k3"]
  fmt.Println("val: ", val, "prs: ", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
}


