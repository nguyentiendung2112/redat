package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :2345")
	if err := http.ListenAndServe(":2345", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
