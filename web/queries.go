package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	RootQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"people": &graphql.Field{
				Type: graphql.NewList(PersonObject),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					session := mongo.SessionFromContext(p.Context)
					if session == nil {
						return nil, errors.New("Session not found in graphql context")
					}
					coll := session.Client().Database("graphql_mongo_prototype").Collection("people")
					cursor, err := coll.Find(p.Context, bson.D{}, nil)
					if err != nil {
						return nil, err
					}
					results := []interface{}{}
					for cursor.Next(p.Context) {
						var result struct {
							ID        string `json:"id" bson:"_id"`
							FirstName string `json:"firstName" bson:"firstName"`
							LastName  string `json:"lastName" bson:"lastName"`
						}
						if err := cursor.Decode(&result); err != nil {
							fmt.Printf("Error occured fetching record %s\n", err.Error())
							continue
						}
						fmt.Printf("Fetched data value = %v\n", result)
						results = append(results, result)
					}
					return results, nil
				},
			},
		},
	})
)
