package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is a test", r.URL.Path[1:])
	fmt.Println("NIce Shit")
	fmt.Printf("KEY: %s", os.Getenv("key1"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serving :8080")
	http.ListenAndServe(":8080", nil)
}
