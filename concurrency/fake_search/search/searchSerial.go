package search

// serial search.
func Serial(query string) ([]Result, error) {
	return []Result{Web1(query), Image1(query), Video1(query)}, nil
}
