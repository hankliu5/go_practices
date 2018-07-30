package search

import (
	"errors"
	"fmt"
	"time"
)

// runs replica search and return the first result from the channel.
func FirstResult(query string, replicas ...SearchFunc) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// replica search to reduce the timeout opportunity
func Replicated(query string, timeout time.Duration) ([]Result, error) {
	c := make(chan Result, 3)
	go func() { c <- FirstResult(query, Web1, Web2) }()
	go func() { c <- FirstResult(query, Image1, Image2) }()
	go func() { c <- FirstResult(query, Video1, Video2) }()
	// timeout := time.After(80 * time.Millisecond)
	timer := time.After(timeout)

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timer:
			fmt.Println("timed out")
			return results, errors.New("timed out")
		}
	}
	return results, nil
}
