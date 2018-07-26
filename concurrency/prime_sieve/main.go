package main

import (
	"fmt"
	"time"
)

// use a goroutine to keep generating number sequence
// channel is synchronous
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// take input number sequence from previous filter, use the prime to keep filter out non-prime number
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	const n = 1000
	start := time.Now()
	ch := make(chan int)
	// the source that generates sequences.
	go generate(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		// fmt.Println(prime)
		ch1 := make(chan int)
		// generate a new goroutine here as a filter that holds prime to filter out not-prime numbers
		go filter(ch, ch1, prime)
		// make the filtered output be a new input for next filter goroutine
		ch = ch1
	}
	fmt.Println(time.Since(start))
}
