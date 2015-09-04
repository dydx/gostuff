package main

import "fmt"

type rect struct {
  width, height int
}

// the `area` method has a reciever of type `rect`
func (r *rect) area() int {
  return r.width * r.height
}

// the `perim` method has a receiver of type `rect`
func (r *rect) perim() int {
  return 2*r.width + 2*r.height
}

func main() {
  r := rect{width: 10, height: 5}

  fmt.Println("Area:", r.area())
  fmt.Println("Perimeter:", r.perim())

  rp := &r
  rp.height = 20
  fmt.Println("Rect:", rp)
  fmt.Println("Area:", rp.area())
  fmt.Println("Perimeter:", rp.perim())
}
