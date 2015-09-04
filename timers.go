package main

import (
  "fmt"
  "time"
)

func main() {

  // timers represent a single event in the future.
  // we tell the timer how long we want to wait, and
  // it provides a channel that will be notified at
  // that time. this timer waits 2 seconds
  timer1 := time.NewTimer(time.Second * 2)


  // this statement blocks on the timer's channel C
  // untl it sends a value indicating that the timer
  // has expired
  <-timer1.C
  fmt.Println("Timer 1 expired")

  // if we just wanted to wait, we could have just
  // used time.Sleep. One reason a timer might be useful
  // is that you can cancel the timer before it expires
  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("Timer 2 expired")
  }()

  // the first timer will expire ~2s after the start,
  // but the second should be stopped before it has a
  // chance to expire
  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("Timer 2 stopped")
  }
}
