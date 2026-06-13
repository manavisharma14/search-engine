package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SearchResult struct {
	ID    string
	Score int
}

var index = map[string][]string{
	"kubernetes":    {"5"},
	"container":     {"5"},
	"microservices": {"6"},
	"scaling":       {"6"},
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	scores := make(map[string]int)
	query := r.URL.Query().Get("q")
	words := strings.Fields(query)

	for _, word := range words {
		ids := index[word]
		for _, id := range ids {
			scores[id]++
		}
	}
	fmt.Println(scores)

	results := []SearchResult{}

	for id, score := range scores {
		results = append(results, SearchResult{
			ID:    id,
			Score: score,
		})
	}

	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/search", searchHandler)
	fmt.Println("shard server running on :5003")
	http.ListenAndServe(":5003", nil)
}
