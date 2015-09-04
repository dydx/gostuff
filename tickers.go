package main

// tickers are for when we want to do something
// repeatedly at regular intervals

import (
  "fmt"
  "time"
)

func main() {

  // tickers use a similar mechanism to timers: a
  // channel that is sent values. Here we'll use
  // the range builtin on the channel to iterate
  // over the values as they arrive ever 500ms
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
    for t := range ticker.C {
      fmt.Println("Tick at", t)
    }
  }()

  // tickers can be stopped just like timers. Once
  // a ticker is stopped, it wont receive any more
  // values on its channel
  time.Sleep(time.Millisecond * 1600)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
