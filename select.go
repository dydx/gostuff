package main

import (
  "time"
  "fmt"
)

func main() {

  // starting with two channels we can select over
  c1 := make(chan string)
  c2 := make(chan string)

  // each channel will receive a value after some amount
  // of time, to similate operations executing in concurrent
  // goroutines
  go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
  }()

  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
  }()

  for i := 0; i < 2; i++ {
    // we'll use select to away both of these values
    // simulteneously, printing each one as it arrives
    select {
    case msg1 := <-c1:
      fmt.Println("received", msg1)
    case msg2 := <-c2:
      fmt.Println("recieved", msg2)
    }
  }
}
