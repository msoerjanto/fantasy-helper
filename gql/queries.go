package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/analytics"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(analyticsService analytics.AnalyticsService) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{analyticsService: analyticsService}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"playerAverages": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(PlayerAverage),
						Args: graphql.FieldConfigArgument{
							"season": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.PlayerAverageResolver,
					},
				},
			},
		),
	}
	return &root
}
