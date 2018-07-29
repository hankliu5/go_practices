package fakesearch

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
