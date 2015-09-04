package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  // here's a non-blocking receive. If a value is
  // available on messages, then select will take
  // <-messages cas with the value, if not, it will
  // immediately take the default case
  select {
  case msg := <-messages:
    fmt.Println("received message:",msg)
  default:
    fmt.Println("no message received")
  }

  // a non-blocking send behaves similarly
  msg := "hi"
  select {
  case messages <- msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }

  // we can use multiple cases above the default
  // clause to implement a multi-way non-blocking
  // select. Here we attempt non-blocking receives
  // on both messages and signals1
  select {
  case msg := <- messages:
    fmt.Println("received message:", msg)
  case sig := <- signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }
}
