package main

import (
	"fmt"
	"time"
)

/*
const sleep = require('util').promisify(setTimeout)
async function myAsyncFunction() {
    await sleep(2000)
    return 2
};

(async function() {
    const result = await myAsyncFunction();
    // outputs `2` after two seconds
    console.log(result);
})();
*/

// async
func myAsyncFunction() <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		time.Sleep(time.Second * 2)
		r <- 2
	}()
	return r
}

func main() {
	r := myAsyncFunction()

	// await
	fmt.Println(<-r)
}
