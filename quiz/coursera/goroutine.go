//I wrote 2 goroutines, which modify the same global variable – x.
//
//The first goroutine, named incrementX10000Times, runs a loop that increments x 10000 times.
//The second goroutine, named decrementX10000Times, runs a loop that decrements x 10000 times.
//I then let these 2 goroutines run concurrently, after initializing x to zero. When done, I print out
//the value of x, after waiting for 100 msec (otherwise the printing will take place before the execution
//of the 2 goroutines is complete and will always print out that x is 0).
//
//The concurrent run of these 2 goroutines is repeated 100 times (which can be modified by adjusting the
//numTests variable). If there was no race condition, the value of x would have always been 0, since the
//number of increments and decrements is identical. However, it’s easy to see that the value of x, after
//the 2 goroutines finish executing, is different in each run and is non-deterministic. This shows that
//we have a race condition in the code.
//
//The reason for the race condition is that both goroutines access the same variable, x, in an unsynchronized manner.

package main

import (
	"fmt"
	"time"
)

var x int

func incrementX10000Times() {
	for i := 0; i < 10000; i++ {
		x = x + 1
	}
}

func decrementX10000Times() {
	for i := 0; i < 10000; i++ {
		x = x - 1
	}
}

func main() {
	numTests := 100
	for i := 0; i < numTests; i++ {
		x = 0
		go incrementX10000Times()
		go decrementX10000Times()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Value of x after test no. ", i+1, "is", x)
	}
}
