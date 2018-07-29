package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := SearchSerial("golang")
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = SearchParallel("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = SearchTimeout("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)

	start = time.Now()
	result := FirstResult("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(result)

	start = time.Now()
	results = SearchReplicated("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)
}
