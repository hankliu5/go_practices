package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hankliu5/go_practices/concurrency/fake_search/search"
)

var timeout = 80 * time.Millisecond

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results, _ := search.Serial("golang")
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results, _ = search.Parallel("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results, _ = search.Timeout("golang", timeout)
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)

	start = time.Now()
	result := search.FirstResult(
		"golang",
		search.FakeSearch("web1", "The Go Programming Language", "http://golang.org"),
		search.FakeSearch("web2", "The Go Programming Language", "http://golang.org"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(result)

	start = time.Now()
	results, _ = search.Replicated("golang", timeout)
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)
}
