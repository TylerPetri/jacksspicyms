package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "<h1>Hello World!</h1>")
	})

	// print a log message
	fmt.Println("Listening on port 8080")

	// start the server
	http.ListenAndServe(":8080", nil)
}
