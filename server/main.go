package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2>Contact Us</h2>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Mom!")
	})
	http.HandleFunc("/contact", contact)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
