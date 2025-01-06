package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

func main() {
	fields := graphql.Fields{
		"User1": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Tahshin", nil
			},
		},
		"User2": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Naheed", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("Faild to create new schema, error: %v", err)
	}
	query := `{
		User1
		User2
		}
`
	params := graphql.Params{Schema: schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation errors: %+v", r.Errors)
	}

	rJson, _ := json.Marshal(r)

	fmt.Println(string(rJson))
}
