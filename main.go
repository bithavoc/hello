package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is a %s\n", os.Getenv("STUFF"))
	fmt.Fprintf(w, "Server")
	fmt.Println("NIce Shit")
	log.Println("error test")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serving :8080")
	http.ListenAndServe(":8080", nil)
}
