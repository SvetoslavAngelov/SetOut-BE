package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SvetoslavAngelov/tourplan-app/src/attraction"
	"github.com/go-chi/chi/v5"
)

var attractions = []attraction.Outline{
	{Id: 1,
		Name:      "MyfirstAttr",
		IsOpen:    false,
		Distance:  2.64,
		Rating:    4.7,
		Latitude:  -0.000536,
		Longitude: 51.476833,
		ImageName: "default"},
}

func getAttraction(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	for _, attr := range attractions {

		// Casting down to int32, unlikely to max it out.
		if attr.Id == int32(id) {
			json.NewEncoder(wr).Encode(attr)
			return
		}

	}

	wr.WriteHeader(http.StatusNotFound)
	json.NewEncoder(wr).Encode("{}")

}

func main() {
	router := chi.NewRouter()
	router.Get("/attraction/{id:[0-9]+}", getAttraction)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("There's an error with the server", err)
	}
}
