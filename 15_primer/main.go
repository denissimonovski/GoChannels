package main

import (
	"sync"
	"fmt"
)

func main() {
	input := gen()
	f1 := fact(input)
	f2 := fact(input)
	f3 := fact(input)
	f4 := fact(input)
	f5 := fact(input)
	var i int
	for n := range merge_fact(f1, f2, f3, f4, f5) {
		i++
		fmt.Println(i, n)
	}
}

func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 2; k < 12; k++ {
					out <- k
				}
			}
		}
		close(out)
	}()
	return out
}

func fact(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func merge_fact(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
