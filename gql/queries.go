package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/analytics"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

var puntArg = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "PuntCategories",
		Fields: graphql.InputObjectConfigFieldMap{
			"points": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"assists": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"rebounds": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"steals": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"blocks": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"fgp": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"ftp": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"threepm": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"to": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(analyticsService analytics.AnalyticsService) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{analyticsService: analyticsService}

	// arguments to the query
	args := graphql.FieldConfigArgument{
		"season": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"puntCategories": &graphql.ArgumentConfig{
			Type: puntArg,
		},
		"includeAll": &graphql.ArgumentConfig{
			Type:         graphql.Boolean,
			DefaultValue: true,
		},
		"numTeams": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 12,
		},
		"rosterSlots": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 14,
		},
	}

	// Create a new Root that describes our base query set up.

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"playerAverages": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type:    graphql.NewList(PlayerAverage),
						Args:    args,
						Resolve: resolver.PlayerAverageResolver,
					},
				},
			},
		),
	}
	return &root
}
