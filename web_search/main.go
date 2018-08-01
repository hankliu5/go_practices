package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/hankliu5/go_practices/concurrency/fake_search/search"
)

var listenAddr = "localhost:4000"
var responseTemplate = template.Must(template.ParseFiles("result.html"))

type response struct {
	Results []search.Result
	Elapsed time.Duration
}

func main() {
	http.HandleFunc("/search", handleSearch)
	fmt.Println("Serving on http://localhost:4000/search")
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
		return
	}
	start := time.Now()
	timeout := 100 * time.Millisecond
	results, err := search.Replicated(query, timeout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	elapsed := time.Since(start)
	resp := response{results, elapsed}
	switch req.FormValue("output") {
	case "json":
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Fatal(err)
		}
	case "prettyjson":
		var b []byte
		b, err := json.MarshalIndent(resp, "", "  ")
		if err == nil {
			_, err = w.Write(b)
		}
	default:
		err := responseTemplate.Execute(w, resp)
		if err != nil {
			log.Fatal(err)
		}
	}
}
