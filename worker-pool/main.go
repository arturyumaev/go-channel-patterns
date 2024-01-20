package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	tasksNum = 100
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(tasksNum)

	tasksCh := make(chan int)
	workers := runtime.NumCPU()

	for w := 0; w < workers; w++ {
		w := w

		go func() {
			for task := range tasksCh {
				fmt.Printf("worker %d has done task %d\n", w, task)
				wg.Done()
			}
		}()
	}

	for t := 0; t < tasksNum; t++ {
		tasksCh <- t
	}
	close(tasksCh)

	wg.Wait()
}
