package main

import (
	"fmt"
	"math/rand"
	"time"
)

// process simulates a process that takes some time to execute.
func process() int {
	time.Sleep(2 * time.Second)
	return 3
}

// futuresDemo demostrates the future (async/awit) pattern in Go.
func futuresDemo() {
	future := make(chan int, 1)
	go func() { future <- process() }() // async

	result := <-future // await
	fmt.Printf("result=%v\n", result)
}

type result struct {
	val int
	err error
}

func rand_process() (int, error) {
	sec := rand.Intn(10)
	time.Sleep(time.Duration(sec) * time.Second)
	return sec, nil
}

func scatterGatherDemo() {

	// Scatter
	c := make(chan result, 10)
	for i := 0; i < cap(c); i++ {
		go func() {
			val, err := rand_process()
			c <- result{val: val, err: err}
		}()
	}

	// Gather
	var total int
	for i := 0; i < cap(c); i++ {
		result := <-c
		fmt.Printf("result=%v\n", result)
		if result.err == nil {
			total += result.val
		}
	}

	fmt.Printf("total=%v\n", total)
}

func main() {
	// futuresDemo()
	// scatterGatherDemo()
}
