/*
https://go.dev/doc/tutorial/getting-started

start the server:
go run main.go

Then visit http://localhost:8080
*/

package main

import (
	"html/template"
	"net/http"
)

type Movie struct {
	Title  string
	Year   int
	Genre  string
	Rating float64
}

func main() {
	// Define a template
	const movieTemplate = `
        <h1>{{ .Title }}</h1>
        <p>Year: {{ .Year }}</p>
        <p>Genre: {{ .Genre }}</p>
        <p>Rating: {{ .Rating }}</p>
    `

	// Parse the template
	tmpl := template.Must(template.New("movie").Parse(movieTemplate))

	// Define some movie data
	movies := []Movie{
		{"Inception", 2010, "Sci-Fi", 8.8},
		{"The Shawshank Redemption", 1994, "Drama", 9.3},
		{"The Matrix", 1999, "Action", 8.7},
	}

	// Create a handler function to render the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render the template with the first movie in the list
		err := tmpl.Execute(w, movies[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
