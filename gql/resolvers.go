package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/analytics"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	analyticsService analytics.AnalyticsService
}

func convertPuntCategories(puntCategories map[string]interface{}) analytics.PuntCategories {
	result := analytics.PuntCategories{}
	if val, ok := puntCategories["points"]; ok {
		result.Points = val.(bool)
	}
	if val, ok := puntCategories["assists"]; ok {
		result.Assists = val.(bool)
	}
	if val, ok := puntCategories["rebounds"]; ok {
		result.TotalRebounds = val.(bool)
	}
	if val, ok := puntCategories["steals"]; ok {
		result.Steals = val.(bool)
	}
	if val, ok := puntCategories["blocks"]; ok {
		result.Blocks = val.(bool)
	}
	if val, ok := puntCategories["fgp"]; ok {
		result.FGPercentage = val.(bool)
	}
	if val, ok := puntCategories["ftp"]; ok {
		result.FTPercentage = val.(bool)
	}
	if val, ok := puntCategories["threepm"]; ok {
		result.ThreePTMade = val.(bool)
	}
	if val, ok := puntCategories["to"]; ok {
		result.Turnovers = val.(bool)
	}
	return result
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) PlayerAverageResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	season, seasonOk := p.Args["season"].(int)
	puntArg, puntArgOk := p.Args["puntCategories"].(map[string]interface{})

	numTeams := p.Args["numTeams"].(int)
	rosterSlots := p.Args["rosterSlots"].(int)
	includeAll := p.Args["includeAll"].(bool)

	if seasonOk {
		var numConsidered int
		if includeAll {
			numConsidered = -1
		} else {
			numConsidered = numTeams * rosterSlots
		}

		var punt analytics.PuntCategories
		if !puntArgOk {
			punt = analytics.PuntCategories{}
		} else {
			punt = convertPuntCategories(puntArg)
		}
		playerAverages := r.analyticsService.GetPlayerAveragesBySeason(
			season, punt, numConsidered)
		return playerAverages, nil
	}

	return nil, nil
}
