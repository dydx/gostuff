package main

import (
  "fmt"
  "time"
)

func worker(done chan bool) {
  fmt.Print("Working...")

  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}

func main() {

  fmt.Println("Starting")
  done := make(chan bool, 1)

  // start a worker goroutine
  // from the worker on the channel
  go worker(done)

  // if this is removed, the program will exit
  // without ever starting the worker
  <-done

  fmt.Println("Finished")
}
