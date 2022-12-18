package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", indexHandler)

	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	var err error = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
