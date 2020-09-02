package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", handlers.Greeting("Code.education Rocks!"))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
