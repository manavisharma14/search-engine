package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	resultChannel := make(chan []SearchResult)
	
	for _, shardURL := range shards {

		go func (url string){

			resp, err := http.Get(url)

			if err != nil {
				fmt.Println(err)
				return
			}

			var results []SearchResult
			json.NewDecoder(resp.Body).Decode(&results)

			resultChannel <- results

		} (shardURL)	
	}

	allResults := []SearchResult{}

	for range shards {
		shardResults := <- resultChannel
		allResults = append(allResults, shardResults...)
	}
	fmt.Println(allResults)
	
}