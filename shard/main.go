package main

import (
	"fmt"
	"net/http"
)

type Document struct{
	ID		string
	Text	string
}

type SearchResult struct{
	ID		string
	Score	int
}

func searchHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	fmt.Fprintf(w, "shard received query: %s", query)
}

func main(){
	http.HandleFunc("/search", searchHandler)
	
	fmt.Println("shard server running :5001")

	http.ListenAndServe(":5001", nil)
}