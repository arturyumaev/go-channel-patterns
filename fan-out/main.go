package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	workers = 2000
)

// CPU bound workload
// Be careful on high RPS
func work(w int) string {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	return fmt.Sprintf("result from worker %d", w)
}

func main() {
	ch := make(chan string, workers)
	defer close(ch)

	for w := 0; w < workers; w++ {
		w := w

		go func() {
			result := work(w)
			ch <- result
		}()
	}

	for workers > 0 {
		result := <-ch
		fmt.Println(result)

		workers--
	}
}
