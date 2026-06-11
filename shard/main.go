package main

import (
	"fmt"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	fmt.Println(w, "shard received query: %s", query)
}

func main(){
	http.handleFunc("/search", searchHandler)
	
	fmt.Println("shard server running :5001")

	http.ListenAndServe(":5001", nil)
}