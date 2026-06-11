package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.handleFunc("/search", searchHandler)
	
	fmt.Println("shard server running :5001")
	
	http.ListenAndServe(":5001", nil)
}