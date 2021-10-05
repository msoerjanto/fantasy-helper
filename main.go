package main

// We import 4 important libraries
// 1. “net/http” to access the core go http functionality
// 2. “fmt” for formatting our text
// 3. “html/template” a library that allows us to interact with our html file.
// 4. "time" - a library for working with date and time.
import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/msoerjanto/fantasy-helper/bballref"
	"github.com/msoerjanto/fantasy-helper/gql"
	"github.com/msoerjanto/fantasy-helper/server"
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
	// Create a new router
	router := chi.NewRouter()

	bballrefService := bballref.NewBasketballRefService()

	// Create our root query for graphql
	rootQuery := gql.NewRoot(bballrefService)
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
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log api request calls
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router
}
