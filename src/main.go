package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Attraction struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	IsOpen   bool    `json:"isopen"`
	Distance float32 `json:"distance"`
	Rating   float32 `json:"rating"`
}

var attractions = []Attraction{
	{Id: 1, Name: "MyfirstAttr", IsOpen: false, Distance: 21.5, Rating: 4.9},
}

func attraction(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	for _, attr := range attractions {
		if attr.Id == id {
			json.NewEncoder(wr).Encode(attr)
			return
		}

	}

	wr.WriteHeader(http.StatusNotFound)
	json.NewEncoder(wr).Encode("{}")

}

func main() {
	router := chi.NewRouter()
	router.Get("/attraction/{id:[0-9]+}", attraction)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("There's an error with the server", err)
	}
}
