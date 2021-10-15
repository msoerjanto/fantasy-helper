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

var DataTableResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DataTableResponse",
		Fields: graphql.Fields{
			"playerAverages": &graphql.Field{
				Type: graphql.NewList(PlayerAverage),
			},
			"aggregates": &graphql.Field{
				Type: Aggregates,
			},
		},
	},
)

var Aggregates = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Aggregates",
		Fields: graphql.Fields{
			"averages": &graphql.Field{
				Type: PlayerAverage,
			},
			"standardDeviations": &graphql.Field{
				Type: PlayerAverage,
			},
		},
	},
)

var LeagueDataResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LeagueData",
		Fields: graphql.Fields{
			"teams": &graphql.Field{
				Type: graphql.NewList(Team),
			},
		},
	},
)

var Manager = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Manager",
		Fields: graphql.Fields{
			"managerId": &graphql.Field{
				Type: graphql.Float,
			},
			"nickname": &graphql.Field{
				Type: graphql.String,
			},
			"guid": &graphql.Field{
				Type: graphql.String,
			},
			"isCurrentLogin": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var Team = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Team",
		Fields: graphql.Fields{
			"managers": &graphql.Field{
				Type: graphql.NewList(Manager),
			},
			"players": &graphql.Field{
				Type: graphql.NewList(PlayerAverage),
			},
		},
	},
)
