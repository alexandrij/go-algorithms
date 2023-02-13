package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func random(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- rand.Intn(n)
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := random(10)

	c1 := sq(in)
	c2 := sq(in)

	for d := range merge(c1, c2) {
		fmt.Println(d)
	}
}
