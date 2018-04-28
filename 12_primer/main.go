package main

import "fmt"

func main() {
	for n := range fact(gen()) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 5; j < 15; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fact(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range c {
			out <- factorial(i)
		}
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
