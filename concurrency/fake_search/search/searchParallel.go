package search

// concurrent search but it still put out the result from channel one by one.
func Parallel(query string) ([]Result, error) {
	c := make(chan Result)

	// run these searches concurrently but these goroutines are waiting for putting the result into the channel.
	go func() { c <- Web1(query) }()
	go func() { c <- Image1(query) }()
	go func() { c <- Video1(query) }()

	return []Result{<-c, <-c, <-c}, nil
}
