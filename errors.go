package main

import (
  "errors"
  "fmt"
)

// by convention, errors are the last return value
// and have type error, a built-in interface

func f1(arg int) (int, error) {
  if arg == 42 {
    // errors.new constructs a basic error value
    // with the given error message
    return -1, errors.New("Can't work with 42")
  }

  // a nil value in the error position indicates
  // that there was no error
  return arg + 3, nil
}

// it is possible to use custom types as errors by
// implementing the Error method on them

// this is a variant on the example above that uses
// a custom type to explicitly represent an argument
// error

type argError struct {
  arg int
  prob string
}

func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
  if arg == 42 {
    // in this case we use &argError syntax to build
    // a new struct, supplying the values for the two
    // fields arg and prob
    return -1, &argError{arg, "can't work with it"}
  }
  return arg + 3, nil
}

func main() {
  // these two loops test out each of the error-returning
  // functions. inline error checking is idiomatic in Go
  for _, i := range []int{7, 42} {
    if r, e := f1(i); e != nil {
      fmt.Println("f1 failed:", e)
    } else {
      fmt.Println("f1 worked:", r)
    }
  }

  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed:", e)
    } else {
      fmt.Println("f2 worked:", r)
    }
  }

  // if you want to programmatically use the data
  // in a custom error, you'll need to get the error
  // as an instance of the custom error via type assertion
  _, e := f2(41)
  if ae, ok := e.(*argError); ok {
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}
