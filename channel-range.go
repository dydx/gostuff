package main

import "fmt"

func main() {

  // iterate over 2 values in the
  // queue channel
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
//  queue <- "three" // this causes a deadlock
  close(queue)

  // this range iterates over each element
  // as it's received from the queue
  // because the chan is closed, iteration
  // terminates after recieving the two
  // elements. if it weren't closed, it'd
  // block on the 3rd receive attempt
  for elem := range queue {
    fmt.Println(elem)
  }
}
