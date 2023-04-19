package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type contextKey string

const neo4jDriverKey contextKey = "neo4jDriver"

func Neo4jDriverMiddleware(driver neo4j.DriverWithContext) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), neo4jDriverKey, driver)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func respondToError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
