package main

import "fmt"

func main() {

  // sum the numbers in a slice
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  // iterate over a slice and operate on the indexes
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index of 3:", i)
    }
  }

  // declare a map and iterate over its contents, printing
  // keys and values
  kvs := map[string]string{"a": "apple", "b": "bananaa"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  // iterate over a string and print indexes and charcodes
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
