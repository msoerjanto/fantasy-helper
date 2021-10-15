package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/analytics"
	"github.com/msoerjanto/fantasy-helper/bballref"
	"github.com/msoerjanto/fantasy-helper/gql"
	"github.com/msoerjanto/fantasy-helper/server"
	"github.com/msoerjanto/fantasy-helper/yahoo"
)

//Go application entrypoint
func main() {

	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	router := initializeAPI()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))

}

func initializeAPI() *chi.Mux {
	// Create a new routerx
	router := chi.NewRouter()

	bballrefService := bballref.NewBasketballRefService()
	analyticsService := analytics.NewAnalyticsService(bballrefService)
	yahooService := yahoo.NewService()

	// Create our root query for graphql
	rootQuery := gql.NewRoot(analyticsService, yahooService)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "X-CSRF-Token", "Authorization"},
		}),
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log api request calls
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())
	router.Get("/auth", yahoo.AuthorizeYahooHandler(yahooService))
	router.Post("/oauth2", yahoo.GetTokenHandler(yahooService))

	return router
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://localhost:3000" {
		return true
	}
	return false
}
