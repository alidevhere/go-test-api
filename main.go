package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Api")
	fmt.Println("Got Hit...")
}
func main() {
	http.HandleFunc("/", Hello)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
