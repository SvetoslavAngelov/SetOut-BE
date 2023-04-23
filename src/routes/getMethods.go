package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SvetoslavAngelov/tourplan-app/src/db_connection"
	"github.com/go-chi/chi/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetAttractionById(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		respondWithError(wr, http.StatusInternalServerError, err)
	}

	// Obtains the driver interface object from the http.Request context
	driver := r.Context().Value(neo4jDriverKey).(neo4j.DriverWithContext)

	// Creates a new session interface, which is then used in the ReadAttractionById function
	// to start a new AuraDB transaction
	session := driver.NewSession(r.Context(), neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(r.Context())

	result, err := db_connection.ReadAttractionById(r.Context(), session, id)

	if err != nil {
		respondWithError(wr, http.StatusInternalServerError, err)
	}

	json.NewEncoder(wr).Encode(result)
}

func GetAttractionsList(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")

	// Obtains the driver interface object from the http.Request context
	driver := r.Context().Value(neo4jDriverKey).(neo4j.DriverWithContext)

	// Creates a new session interface, which is then used in the ReadAttractionById function
	// to start a new AuraDB transaction
	session := driver.NewSession(r.Context(), neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(r.Context())

	results, err := db_connection.ReadAttractions(r.Context(), session)

	if err != nil {
		respondWithError(wr, http.StatusInternalServerError, err)
	}

	json.NewEncoder(wr).Encode(results)
}
