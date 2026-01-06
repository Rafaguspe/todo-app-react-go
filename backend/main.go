package main

import (
	"fmt"
	"log"
	"net/http"
)

func ResponseOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Init server")
}

func main() {
	http.HandleFunc("/", ResponseOne)

	fmt.Println("Server init on http://localhost:8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
