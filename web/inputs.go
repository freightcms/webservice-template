package web

import (
	"errors"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
					if err := session.StartTransaction(); err != nil {
						return nil, err
					}
					createPersonModel := struct {
						FirstName string `json:"firstName" bson:"firstName"`
						LastName  string `json:"lastName" bson:"lastName"`
					}{
						FirstName: params.Args["firstName"].(string),
						LastName:  params.Args["lastName"].(string),
					}
					coll := session.Client().Database("graphql_mongo_prototype").Collection("people")
					insertedResult, err := coll.InsertOne(params.Context, &createPersonModel, options.InsertOne())
					if err != nil {
						return nil, err
					}
					session.CommitTransaction(params.Context)
					var result struct {
						ID        string `json:"id" bson:"_id"`
						FirstName string `json:"firstName" bson:"firstName"`
						LastName  string `json:"lastName" bson:"lastName"`
					}

					filter := bson.D{{"_id", insertedResult.InsertedID.(primitive.ObjectID).String()}}
					if err := coll.FindOne(params.Context, filter).Decode(&result); err != nil {
						return nil, err
					}

					return result, nil
				},
			},
		},
	})
)
