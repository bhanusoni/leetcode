# Go Concurrency Patterns

This note covers practical patterns you can reuse in interviews and production code.

## 1) Worker Pool

Use when you have many independent jobs and want bounded parallelism.

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * j
		fmt.Printf("worker %d processed %d\n", id, j)
	}
}

func main() {
	const workerCount = 3
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("result:", r)
	}
}
```

## 2) Fan-Out / Fan-In

Use fan-out to split work across goroutines and fan-in to merge outputs.

```go
func fanOut(in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, 0, n)
	for i := 0; i < n; i++ {
		out := make(chan int)
		go func(ch chan<- int) {
			defer close(ch)
			for v := range in {
				ch <- v * 2
			}
		}(out)
		outs = append(outs, out)
	}
	return outs
}

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

## 3) Pipeline With Cancellation (`context.Context`)

Use when multi-stage processing should stop early on timeout/error.

```go
func stage(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v + 1:
			}
		}
	}()
	return out
}
```

## 4) Semaphore (Limit Concurrent Calls)

Use a buffered channel as a lightweight semaphore.

```go
sem := make(chan struct{}, 5) // max 5 concurrent tasks
var wg sync.WaitGroup

for _, task := range tasks {
	wg.Add(1)
	go func(t Task) {
		defer wg.Done()
		sem <- struct{}{}        // acquire
		defer func() { <-sem }() // release
		process(t)
	}(task)
}

wg.Wait()
```

## 5) Timeout With `select`

Use when a channel receive should not block forever.

```go
select {
case msg := <-ch:
	fmt.Println("received:", msg)
case <-time.After(2 * time.Second):
	fmt.Println("timed out")
}
```

## Common Mistakes To Avoid

- Forgetting to close channels owned by a producer.
- Closing a channel from the receiver side.
- Starting goroutines without a stop path (`context`, done channel, or bounded input).
- Writing to an unbuffered channel without a receiver ready.
- Ignoring race conditions (run with `go test -race`).

