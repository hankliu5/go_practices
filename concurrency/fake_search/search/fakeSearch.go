package search

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = FakeSearch("web1", "The Go Programming Language", "http://golang.org")
	Web2   = FakeSearch("web2", "The Go Programming Language", "http://golang.org")
	Image1 = FakeSearch("image1", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Image2 = FakeSearch("image2", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video1 = FakeSearch("video1", "Concurrency is not Parallelism",
		"https://www.youtube.com/watch?v=cN_DpYBzKso")
	Video2 = FakeSearch("video2", "Concurrency is not Parallelism",
		"https://www.youtube.com/watch?v=cN_DpYBzKso")
)

type SearchFunc func(query string) Result

type Result struct {
	Title, URL string
}

// the main search function
func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s result for %q: %s\n", kind, query, title),
			URL:   url,
		}
	}
}
