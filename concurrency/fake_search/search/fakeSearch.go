package search

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = FakeSearch("web")
	Image  = FakeSearch("image")
	Video  = FakeSearch("video")
	Web2   = FakeSearch("web")
	Image2 = FakeSearch("image")
	Video2 = FakeSearch("video")
)

type Search func(query string) Result
type Result string

// the main search function
func FakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
