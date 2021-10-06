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

			"astZ": &graphql.Field{
				Type: graphql.Float,
			},
			"ptZ": &graphql.Field{
				Type: graphql.Float,
			},
			"rebZ": &graphql.Field{
				Type: graphql.Float,
			},
			"stlZ": &graphql.Field{
				Type: graphql.Float,
			},
			"blkZ": &graphql.Field{
				Type: graphql.Float,
			},
			"fgpZ": &graphql.Field{
				Type: graphql.Float,
			},
			"ftpZ": &graphql.Field{
				Type: graphql.Float,
			},
			"toZ": &graphql.Field{
				Type: graphql.Float,
			},
			"threeZ": &graphql.Field{
				Type: graphql.Float,
			},
			"zscore": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
