package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = fakeSearch("web")
	Image  = fakeSearch("image")
	Video  = fakeSearch("video")
	Web2   = fakeSearch("web")
	Image2 = fakeSearch("image")
	Video2 = fakeSearch("video")
)

type Search func(query string) Result

type Result string

// the main search function
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// sequence search.
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// concurrent search but it still put out the result from channel one by one.
func Google2_0(query string) (results []Result) {
	c := make(chan Result)

	// run these searches concurrently but these goroutines are waiting for putting the result into the channel.
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

// concurrent search with timeout so it doesn't wait for the long responses.
func Google2_1(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// runs replica search and return the first result from the channel.
func FirstResult(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// replica search to reduce the timeout opportunity
func Google3_0(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- FirstResult(query, Web, Web2) }()
	go func() { c <- FirstResult(query, Image, Image2) }()
	go func() { c <- FirstResult(query, Video, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = Google2_0("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	results = Google2_1("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)

	start = time.Now()
	result := FirstResult("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(result)

	start = time.Now()
	results = Google3_0("golang")
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(results)
}
