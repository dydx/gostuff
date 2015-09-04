package main

import "fmt"

// int parameters, arguments will be passed
// by value. zeroval will get a copy of ival
// distinct from the one in the calling function
func zeroval(ival int) {
  ival = 0
}

// takes an int pointer as its parameter.
// *iptr defreferences the pointer from its
// address to the current value at that address
// assigning a value to a dereferenced pointer
// changes the value at the referenced address
func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1

  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  // this syntax gives the memory address of i
  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  // zeroval doesnt change i in main, but zeroptr does
  // because it has a reference to the memory address
  // for that variable
  fmt.Println("pointer:", &i)
}
