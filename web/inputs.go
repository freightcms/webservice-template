package web

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/webservice-prototype/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createPerson": &graphql.Field{
				Type:        PersonObject,
				Description: "Create new Person",
				Args: graphql.FieldConfigArgument{
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					session := mongo.SessionFromContext(params.Context)
					if session == nil {
						return nil, errors.New("failed to fetch session from context")
					}
					createPersonModel := struct {
						FirstName string `json:"firstName" bson:"firstName"`
						LastName  string `json:"lastName" bson:"lastName"`
					}{
						FirstName: params.Args["firstName"].(string),
						LastName:  params.Args["lastName"].(string),
					}
					coll := session.Client().Database("graphql_mongo_prototype").Collection("people")
					insertedResult, err := coll.InsertOne(
						params.Context,
						&createPersonModel,
						options.InsertOne(),
					)
					if err != nil {
						return nil, err
					}

					var result models.Person
					filter := bson.M{"_id": insertedResult.InsertedID}
					if err := coll.FindOne(params.Context, filter).Decode(&result); err != nil {
						return nil, err
					}

					return result, nil
				},
			},
		},
	})
)
