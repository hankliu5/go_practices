package fakesearch

import (
	"fmt"
	"time"
)

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
func SearchReplicated(query string) (results []Result) {
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
