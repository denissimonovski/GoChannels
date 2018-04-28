package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	for i := range merge(c1, c2) {
		fmt.Println(i)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, i := range nums {
			out <- i
		}
		close(out)
	}()

	return out
}

func sq(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(c ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, cs := range c {
		go func(c1 chan int) {
			for i := range c1 {
				out <- i
			}
			wg.Done()
		}(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
