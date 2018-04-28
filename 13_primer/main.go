package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Foo:"), boring("Bar:"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring.\nI'm leaving")
}

func boring(msg string) <-chan string {
	out := make(chan string)

	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

func fanIn(c1, c2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for {
			out <- <-c1
		}
	}()

	go func() {
		for {
			out <- <-c2
		}
	}()
	return out
}
