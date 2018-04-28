package main

import "fmt"

func main() {
	for n := range factorial(4) {
		fmt.Println("Total:", n)
	}
}

func factorial(a int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := a; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
