package main

import (
	"fmt"
	"time"
)

func asyncFunction(s int) <-chan int {
	r := make(chan int)
	go func() {
		defer close(r)

		time.Sleep(time.Second * time.Duration(s))
		r <- s
	}()
	return r
}

func main() {
	firstCh, secondCh := asyncFunction(2), asyncFunction(3)
	first, second := <-firstCh, <-secondCh

	fmt.Println(first, second)
}
