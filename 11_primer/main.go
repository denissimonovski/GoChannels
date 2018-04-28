package main

import "fmt"

func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
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
