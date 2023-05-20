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
func ReadAttractionById(ctx context.Context, session neo4j.SessionWithContext, id int) (attraction.Outline, error) {

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		attractionNode := attraction.MakeOutline()

		txQuery := `
				MATCH (a: Attractions)
				WHERE a.id = $id
				RETURN a
			`
		result, err := tx.Run(ctx, txQuery, map[string]any{
			"id": id,
		})

		// Return an empty object if the query fails.
		if err != nil {
			return attractionNode, err
		}

		if result.Next(ctx) {
			node := result.Record().Values[0].(neo4j.Node)
			attractionNode.Id = node.Props["id"].(int64)
			attractionNode.Name = node.Props["name"].(string)
			attractionNode.Rating = node.Props["rating"].(float64)
			attractionNode.Latitude = node.Props["latitude"].(float64)
			attractionNode.Longitude = node.Props["longitude"].(float64)
			attractionNode.ImageName = node.Props["attractionImageName"].(string)
		}

		return attractionNode, nil
	})

	if err != nil {
		return attraction.MakeOutline(), err
	}

	return result.(attraction.Outline), nil
}

// Retrieve all attraction nodes.
// TODO, a good idea would be to limit the number of nodes to something reasonable, like 5
func ReadAttractions(ctx context.Context, session neo4j.SessionWithContext) ([]attraction.Outline, error) {
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		attractionList := []attraction.Outline{}

		txQuery := `
				MATCH (a: Attractions)
				RETURN a
			`
		result, err := tx.Run(ctx, txQuery, nil)

		// Return an empty list of objects
		if err != nil {
			return attractionList, err
		}

		for result.Next(ctx) {
			node := result.Record().Values[0].(neo4j.Node)
			attractionNode := attraction.MakeOutline()
			attractionNode.Id = node.Props["id"].(int64)
			attractionNode.Name = node.Props["name"].(string)
			attractionNode.Rating = node.Props["rating"].(float64)
			attractionNode.Latitude = node.Props["latitude"].(float64)
			attractionNode.Longitude = node.Props["longitude"].(float64)
			attractionNode.ImageName = node.Props["attractionImageName"].(string)

			attractionList = append(attractionList, attractionNode)
		}

		return attractionList, nil
	})

	if err != nil {
		return []attraction.Outline{}, err
	}

	return result.([]attraction.Outline), nil
}
