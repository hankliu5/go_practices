package search

// concurrent search but it still put out the result from channel one by one.
func SearchParallel(query string) (results []Result) {
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
