package main

import (
	"fmt"
	"math/rand"
)

func gen(ch chan<- int, limit int) {
	defer close(ch)
	for i := 0; i < limit; i++ {
		n := rand.Intn(limit)
		ch <- n
		fmt.Printf("gen: %d\n", n)
	}
}

func sq(numsCh <-chan int, sqCh chan<- int) {
	defer close(sqCh)
	for v := range numsCh {
		n := v * v
		sqCh <- n
		fmt.Printf("sq: %d\n", n)
	}
}

func pipeline() {
	numsCh := make(chan int)
	sqCh := make(chan int)

	go gen(numsCh, 10)
	go sq(numsCh, sqCh)

	numsCh <- 1000

	for v := range sqCh {
		fmt.Println(v)
	}
}

func main() {
	pipeline()
}
