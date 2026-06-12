package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type SearchResult struct {
	ID    string
	Score int
}

func main(){

	shards := []string{
		"http://localhost:5001/search?q=grpc",
		"http://localhost:5002/search?q=redis",
	}
	
	for _, shardURL := range shards {
		resp, err := http.Get(shardURL)

		if err != nil {
			fmt.Println(err)
			continue
		}

		var results []SearchResult
		json.NewDecoder(resp.Body).Decode(&results)
		fmt.Println(results)
	}

	
}