package db_connection

import (
	"context"

	"github.com/SvetoslavAngelov/tourplan-app/src/attraction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

/*
	A collection of helper functions to retrieve data from AuraDB
*/

// Retrieve a single tourist attraction node, given an AuraDB session and an attraction ID.
func ReadAttractionById(session neo4j.SessionWithContext, id int32) (attraction.Outline, error) {

	ctx := context.Background()
	location := attraction.MakeOutline()

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		readLocationById := `
				MATCH (a: Attractions)
				WHERE a.id = $id
				return a
			`
		result, err := tx.Run(ctx, readLocationById, map[string]any{
			"locationId": id,
		})

		// Return an empty object if the query fails.
		if err != nil {
			return location, result.Err()
		}

		if result.Next(ctx) {
			record := result.Record()
			locationNode := record.Values[0].(neo4j.Node)
			location.Id = locationNode.Props["id"].(int32)
			location.Name = locationNode.Props["name"].(string)
			location.Rating = locationNode.Props["rating"].(float32)
			location.Latitude = locationNode.Props["latitude"].(float64)
			location.Longitude = locationNode.Props["longitude"].(float64)
			location.ImageName = locationNode.Props["attractionImageName"].(string)
		}

		return location, result.Err()
	})

	if err != nil {
		return location, err
	}

	return location, nil
}
