package main

import (
	"fmt"
	"time"
)

func generate(data []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, v := range data {
			out <- v
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return out
}

func inc(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			out <- v + 1
		}
	}()

	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func main() {
	data := generate([]int{1, 2, 3, 4, 5})

	result := sqr(inc(data))

	for v := range result {
		fmt.Println(v)
	}
}
