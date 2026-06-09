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
<!DOCTYPE html>
<html>
<head>
    <title>Go Search Engine</title>
    <style>
body {
    font-family: Inter, sans-serif;
    background: #FAFAFA;
    padding: 60px;
}

.container {
    max-width: 900px;
    margin: auto;
}

h1 {
    text-align: center;
    font-size: 52px;
    color: #111111;
    margin-bottom: 8px;
    font-weight: 700;
}

.subtitle {
    text-align: center;
    color: #6B7280;
    margin-bottom: 40px;
}

form {
    display: flex;
    gap: 12px;
    margin-bottom: 30px;
}

input {
    flex: 1;
    padding: 18px;
    border-radius: 18px;
    border: 1px solid #ECECEC;
    background: white;
    font-size: 16px;
}

input:focus {
    outline: none;
    border-color: #FF5A5F;
}

button {
    border: none;
    background: #FF5A5F;
    color: white;
    padding: 18px 26px;
    border-radius: 18px;
    cursor: pointer;
    font-weight: 600;
}

.result {
    background: white;
    border: 1px solid #ECECEC;
    border-radius: 20px;
    padding: 22px;
    margin-bottom: 14px;
    transition: 0.2s ease;
}

.result:hover {
    transform: translateY(-2px);
}

.score {
    margin-top: 8px;
    color: #6B7280;
}
    </style>
</head>

<body>
    <div class="container">
        <h1>🔍 Go Search Engine</h1>

        <form>
            <input name="q" placeholder="Search documents...">
            <button type="submit">Search</button>
        </form>

`)

	for _, result := range results {
		for _, doc := range docs {
			if doc.ID == result.ID {
				fmt.Fprintf(w,
					`<div class="result">
						<div>%s</div>
						<div class="score">Score: %d</div>
					</div>`,
					doc.Text,
					result.Score,
				)
			}
		}
	}
fmt.Fprintln(w, `
    </div>
</body>
</html>
`)
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