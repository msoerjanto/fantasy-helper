package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/bballref"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	bballrefService bballref.BasketballRefService
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) PlayerAverageResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	season, ok := p.Args["season"].(int)
	if ok {
		users := r.bballrefService.GetPlayerAveragesBySeason(season)
		return users, nil
	}

	return nil, nil
}
