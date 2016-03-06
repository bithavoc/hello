package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is a %s", os.Getenv("STUFF"))
	fmt.Println("NIce Shit")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serving :8080")
	http.ListenAndServe(":8080", nil)
}
