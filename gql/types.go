package gql

import "github.com/graphql-go/graphql"

// User describes a graphql object containing a PlayerAverage
var PlayerAverage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlayerAverage",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"team": &graphql.Field{
				Type: graphql.String,
			},
			"gamesPlayed": &graphql.Field{
				Type: graphql.Int,
			},
			"gamesStarted": &graphql.Field{
				Type: graphql.Int,
			},
			"minutesPerGame": &graphql.Field{
				Type: graphql.Float,
			},
			"personalFouls": &graphql.Field{
				Type: graphql.Float,
			},
			"fgMade": &graphql.Field{
				Type: graphql.Float,
			},
			"fgAttempted": &graphql.Field{
				Type: graphql.Float,
			},
			"ftMade": &graphql.Field{
				Type: graphql.Float,
			},
			"ftAttempted": &graphql.Field{
				Type: graphql.Float,
			},
			"fgPercentage": &graphql.Field{
				Type: graphql.Float,
			},
			"ftPercentage": &graphql.Field{
				Type: graphql.Float,
			},
			"threePtMade": &graphql.Field{
				Type: graphql.Float,
			},
			"totalRebounds": &graphql.Field{
				Type: graphql.Float,
			},
			"assists": &graphql.Field{
				Type: graphql.Float,
			},
			"steals": &graphql.Field{
				Type: graphql.Float,
			},
			"blocks": &graphql.Field{
				Type: graphql.Float,
			},
			"turnovers": &graphql.Field{
				Type: graphql.Float,
			},
			"points": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
