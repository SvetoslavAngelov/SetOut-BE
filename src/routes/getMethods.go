package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SvetoslavAngelov/tourplan-app/src/testdata"
	"github.com/go-chi/chi/v5"
)

func GetAttractionById(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(wr, http.StatusText(400), 400)
	}

	for _, object := range testdata.Attractions {
		if object.Id == int32(id) {
			json.NewEncoder(wr).Encode(object)
			return
		}
	}

	wr.WriteHeader(http.StatusNotFound)
	json.NewEncoder(wr).Encode("{}")
}

func GetAttractionsList(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")

	json.NewEncoder(wr).Encode(testdata.Attractions)
}
