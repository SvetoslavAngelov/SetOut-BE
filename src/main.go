package main

import (
	"context"
	"net/http"

	"github.com/SvetoslavAngelov/tourplan-app/src/db_connection"
	"github.com/SvetoslavAngelov/tourplan-app/src/routes"
	"github.com/go-chi/chi/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {

	// Create a new a AuraDB driver & session
	uri := db_connection.AuraDbUri()
	auth := db_connection.AuraDbAuthToken()

	driver, db_err := neo4j.NewDriverWithContext(uri, auth)

	if db_err != nil {
		println("Couldn't create a new database connection, with an error, ", db_err)
	}

	ctx := context.Background()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "RouteData"})
	defer session.Close(ctx)

	// Create a new Router interface
	router := chi.NewRouter()
	router.Use(routes.Neo4jSessionMiddleware(session))

	router.Route("/", func(r chi.Router) {
		r.Get("/attractions/{id:[0-9]+}", routes.GetAttractionById)
		r.Get("/attractions", routes.GetAttractionsList)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("There's an error with the server", err)
	}
}
