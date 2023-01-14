package main

import (
	"net/http"

	"github.com/SvetoslavAngelov/tourplan-app/src/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/attractions/{id:[0-9]+}", routes.GetAttractionById)
		r.Get("/attractions", routes.GetAttractionsList)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("There's an error with the server", err)
	}
}
