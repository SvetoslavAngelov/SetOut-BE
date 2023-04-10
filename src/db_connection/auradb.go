package db_connection

import (
	"context"

	"github.com/SvetoslavAngelov/tourplan-app/src/attraction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Helper class to manage AuraDB resources
type AuraDbHandler struct {
	driver neo4j.Driver
}

// Create new AuraDB handler
func NewAuraDbHandler(uri, username, password string) (*AuraDbHandler, error) {
	auth := neo4j.BasicAuth(username, password, "")
	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		return nil, err
	}

	return &AuraDbHandler{driver: driver}, nil
}

func (a *AuraDbHandler) Close(ctx context.Context) {
	a.driver.Close(ctx)
}

func (a *AuraDbHandler) ReadLocationsById(locationId int) error {
	ctx := context.Background()
	session := a.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "RouteData"})
	defer session.Close(ctx)

	location, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		readLocationById := `
				MATCH (a: Attractions)
				WHERE id(a) = $locationId
				return a
			`
		result, err := tx.Run(ctx, readLocationById, map[string]interface{}{
			"locationId": locationId,
		})

		if err != nil {
			return attraction.Outline{}, result.Err()
		}

		location := attraction.MakeOutline()

		if result.Next(ctx) {
			record := result.Record()
			locationNode := record.Values[0].(neo4j.Node)
			location.Id = locationNode.Id
			location.Name = locationNode.Props["name"].(string)
			location.Rating = locationNode.Props["rating"].(float32)
			location.Latitude = locationNode.Props["latitude"].(float64)
			location.Longitude = locationNode.Props["longitude"].(float64)
			location.ImageName = locationNode.Props["attractionImageName"].(string)
		}

		return location, result.Err()
	})

	if err != nil {
		return attraction.MakeOutline(), err
	}

	return location, err
}
