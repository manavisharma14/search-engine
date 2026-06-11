package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	resp, err := http.Get(
		"http://localhost:5001/search?q=grpc",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}