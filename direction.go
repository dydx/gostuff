package main

import "fmt"

// the ping function only accepts a channel for sending
// values. It would be a compile-time error to try and
// receive on that channel
func ping(pings chan<- string, msg string) {
  fmt.Println("Ping!")
  pings <- msg
}

// the pong function accepts one channel for receives(pings)
// and a second one for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
  fmt.Println("Pong!")
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)

  ping(pings, "passed message")
  pong(pings, pongs)

  fmt.Println(<-pongs)
}
