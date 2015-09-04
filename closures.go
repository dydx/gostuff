package main

import "fmt"

// this is a function that returns another function
// which we define anonymously in the body. the returned
// function closes over the variable i to form a closure
func intSeq() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func main() {

  // we can call intSeq and assign the result
  // to nextInt. This function captures its own i
  // vale which will be updated each time we call
  // nextInt
  nextInt := intSeq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // to confirm that the state is unique to that
  // particular function, create and test a new one
  newInts := intSeq()
  fmt.Println(newInts())
}
