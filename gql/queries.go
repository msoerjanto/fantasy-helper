package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/analytics"
	"github.com/msoerjanto/fantasy-helper/yahoo"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

var tokenArg = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Token",
		Fields: graphql.InputObjectConfigFieldMap{
			"access_token": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"refresh_token": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"expiry": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"token_type": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

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
func NewRoot(analyticsService analytics.AnalyticsService,
	yahooService yahoo.YahooService) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{analyticsService: analyticsService, yahooService: yahooService}

	// arguments to the query
	dataTableArgs := graphql.FieldConfigArgument{
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

	teamsArgs := graphql.FieldConfigArgument{
		"league": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: tokenArg,
		},
	}

	// Create a new Root that describes our base query set up.

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"dataTable": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type:    DataTableResponse,
						Args:    dataTableArgs,
						Resolve: resolver.PlayerAverageResolver,
					},
					"teams": &graphql.Field{
						Type:    graphql.NewList(Team),
						Args:    teamsArgs,
						Resolve: resolver.TeamsResolver,
					},
				},
			},
		),
	}
	return &root
}
