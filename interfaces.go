package main

import (
  "fmt"
  "math"
)

// define a `geometry` interface
type geometry interface {
  area() float64
  perim() float64
}

// define some structs
type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

// define some methods that work
// with the `geometry` interface
func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

// define a function that takes in methods
// that implement the `geometry` interface
func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

// yeah...
func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}

  measure(r)
  measure(c)
}
