package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {
  fmt.Println(person{"Bob", 20})

  fmt.Println(person{name: "Alice", age: 30})

  fmt.Println(person{name: "Fred"})

  fmt.Println(&person{name: "Ann", age: 40})

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sptr := &s
  fmt.Println(sptr.age)

  sptr.age = 51
  fmt.Println(sptr.age)
  fmt.Println(s.age)

  scpy := s
  scpy.age = 52
  fmt.Println(scpy.age)
}
