package main

import ("fmt"
		"net/http"
		"strings"
	)

var index = make(map[string][]string)
var docs []Document

func helloHandle(w http.ResponseWriter, r *http.Request){


	// query := r.URL.Query().Get("q")
	// results := rankResults(index, query)

	// for _, result := range results {
	// 	for _, doc := range docs {
	// 		if doc.ID == result.ID {
	// 			fmt.Fprintln(w, doc.Text, result.Score)
	// 		}
	// 	}
	// }

	query := r.URL.Query().Get("q")
	results := rankResults(index, query)

	fmt.Fprintln(w, `
		<html>
			<body>
				<h1>Search engine</h1>
				<form>
					<input name="q">
					<button>Search</button>
				</form>
			</body>
		</html>
	`)

	for _, result := range results {
	for _, doc := range docs {
		if doc.ID == result.ID {
			fmt.Fprintln(w, doc.Text, result.Score)
		}
	}
}
}

func main(){

	docs = []Document{
		{ID: "1", Text: "go go go is fast"},
		{ID: "2", Text: "go is simple"},
		{ID: "3", Text: "search engines use indexes"},
	}

	for _, doc := range docs{
		words := strings.Fields(doc.Text)

		for _, word := range words {
			index[word] = append(index[word], doc.ID)
		}
	}

	http.HandleFunc("/", helloHandle)
	
	fmt.Println("server running on :8080")

	http.ListenAndServe(":8080", nil)
}