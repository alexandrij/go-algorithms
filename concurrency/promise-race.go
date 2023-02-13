package main

import (
	"fmt"
	"time"
)

func asyncFunc(s int) <-chan int {
	r := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		r <- s
	}()
	return r
}

func main() {
	var r int

	select {
	case r = <-asyncFunc(3):
		break
	case r = <-asyncFunc(2):
		break
	}

	fmt.Println(r)
}
