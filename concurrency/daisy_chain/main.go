package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func run_daisy_chain(n int) {
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}

func run_normal_increment(n int) {
	i := 1
	for ; i <= n; i++ {
	}
	fmt.Println(i)
}

func main() {
	const n = 10000000
	// slower because it needs to create new goroutines, channels and communicate with each other.
	run_daisy_chain(n)

	// faster in this case
	run_normal_increment(n)
}
