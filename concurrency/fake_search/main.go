package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hankliu5/go_practices/concurrency/fake_search/search"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := search.SearchSerial("golang")
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = search.SearchParallel("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = search.SearchTimeout("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)

	start = time.Now()
	result := search.FirstResult("golang", search.FakeSearch("replica 1"), search.FakeSearch("replica 2"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(result)

	start = time.Now()
	results = search.SearchReplicated("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)
}
