package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SvetoslavAngelov/tourplan-app/src/db_connection"
	"github.com/SvetoslavAngelov/tourplan-app/src/testdata"
	"github.com/go-chi/chi/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetAttractionById(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(wr, http.StatusText(400), 400)
	}

	// Obtains the driver interface object from the http.Request context
	driver := r.Context().Value("neo4jDriver").(neo4j.DriverWithContext)

	// Creates a new session interface, which is then used in the ReadAttractionById function
	// to start a new AuraDB transaction
	session := driver.NewSession(r.Context(), neo4j.SessionConfig{DatabaseName: "RouteData"})
	defer session.Close(r.Context())

	// Read a single attraction from the database and encode the result into a JSON
	result, e := db_connection.ReadAttractionById(session, int32(id))
	if e != nil {
		respondToError(wr, http.StatusInternalServerError, err)
		return
	}

	e = json.NewEncoder(wr).Encode(result)

	if e != nil {
		wr.WriteHeader(http.StatusNotFound)
		json.NewEncoder(wr).Encode("{}")
		return
	}
}

func GetAttractionsList(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")

	json.NewEncoder(wr).Encode(testdata.Attractions)
}
