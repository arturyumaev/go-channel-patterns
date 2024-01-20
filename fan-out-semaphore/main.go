package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var (
	tasks   = 500
	workers = runtime.GOMAXPROCS(0)
	acquire = struct{}{}
)

func main() {
	ch := make(chan string, tasks)
	defer close(ch)

	sem := make(chan struct{}, workers)
	defer close(sem)

	for task := 0; task < tasks; task++ {
		task := task

		go func() {
			sem <- acquire

			t := time.Duration(rand.Intn(200)) * time.Millisecond
			time.Sleep(t)
			ch <- fmt.Sprintf("task %d done", task)

			<-sem
		}()
	}

	for tasks > 0 {
		data := <-ch
		fmt.Println(data)

		tasks--
	}
}
