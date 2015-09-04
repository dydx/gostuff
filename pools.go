package main

import (
  "fmt"
  "time"
)

// this is our worker, of which we'll run several
// concurrent instances. these workers will receive
// work on the jobs channel and send the corresponding
// results on results. We'll sleep a second per job
// to simulate an expensive task

func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
}

func main() {

  // in order to use our workers we need to send them
  // work and collect their results. we make 2 channels
  // for this
  jobs := make(chan int, 100)
  results := make(chan int, 100)

  // this starts up three workers, initially blocked
  // because there are no jobs yet
  for w := 1; w <= 5; w++ {
    go worker(w, jobs, results)
  }

  // here we send 9 jobs and then close that channel
  // indicate that's all the work we have
  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)

  // finally we collect all of the results of the work
  for a := 1; a <= 9; a++ {
    <-results
  }

  // This program shows the 9 jobs being executed
  // by various workrs. it only takes about 3 seconds
  // though, because there are three workers operating
  // concurrently
}
