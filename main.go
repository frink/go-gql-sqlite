package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/corganfuzz/go-gql-sqlite/pkg/model"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Println("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

var agregateSchema = graphql.Fields{
	"tutorial": model.SingleTutorialSchema(),
	"list":     model.ListTutorialSchema(),
}

var agregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": model.CreateTutorialMutation(),
	},
})

func main() {
	// db, err := sql.Open("sqlite3", "./tutorials.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: agregateSchema}
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: agregateMutations,
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		mutation {
			create(id: 3, title: "Acid") {
				id
				title
			}
		}
	`
	// params := graphql.Params{Schema: schema, RequestString: query}
	// r := graphql.Do(params)
	// if len(r.Errors) > 0 {
	// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// }
	// rJSON, _ := json.Marshal(r)

	result := executeQuery(query, schema)
	rJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", rJSON)

}
