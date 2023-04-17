package routes

import (
	"context"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type contextKey string

const neo4jSessionKey contextKey = "neo4jSession"

func Neo4jSessionMiddleware(session neo4j.SessionWithContext) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), neo4jSessionKey, session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
