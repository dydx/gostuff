package main

import "fmt"

func main() {

  // creating a jobs channel and a channel to block
  // program execution until all jobs are done

  // there's an interesting thing to observe when
  // pushing more than the '3' jobs in the tut:
  // it flushes the channel or something every 5
  // scheduled jobs
  jobs := make(chan int, 1000)
  done := make(chan bool)

  // repeatedly serves receives jobs over the
  // `jobs` channel and notifies when its done
  go func() {
    for {
      j, more := <-jobs
      if more {
        fmt.Println("received job", j)
      } else {
        fmt.Println(" -> received all jobs")
        done <- true
        return
      }
    }
  }()

  // send a handful of jobs out and close the channel
  // for sends
  for j := 1; j <= 1000000; j++ {
    jobs <- j
    fmt.Println("sent job", j)
  }
  close(jobs)
  fmt.Println(" -> sent all jobs")

  <-done
}
