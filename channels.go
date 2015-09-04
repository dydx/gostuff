package main

import "fmt"

func main() {
  // create a new channel with `make(chan val-type)`
  // channels are typed by the values they convey
  messages := make(chan string)

  // send a value into a channel using <-
  // here we send "ping" to the messages
  // chanel we made above, from a new goroutine
  go func() { messages <- "ping" }()

  // receive a value from the channel
  // here we'll receive the "ping" message we
  // sent above and print it out
  msg := <-messages
  fmt.Println(msg)

  // sends and receives block until both the sender
  // and the receiver are ready. This allows us to wait
  // at the end of the program without needing to set up
  // any special syncronizations
}
