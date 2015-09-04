package main

import "fmt"

func main() {
  // make a channel with a buffer of size 2
  messages := make(chan string, 2)

  messages <- "buffered"
  messages <- "channel"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
