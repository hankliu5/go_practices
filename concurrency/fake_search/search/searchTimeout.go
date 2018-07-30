package search

import (
	"errors"
	"fmt"
	"time"
)

// concurrent search with timeout so it doesn't wait for the long responses.
func Timeout(query string, timeout time.Duration) ([]Result, error) {
	c := make(chan Result, 3)
	go func() { c <- Web1(query) }()
	go func() { c <- Image1(query) }()
	go func() { c <- Video1(query) }()

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
