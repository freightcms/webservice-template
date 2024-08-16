package web

import (
	"github.com/graphql-go/graphql"
	"github.com/squishedfox/webservice-prototype/db/mongodb"
	"github.com/squishedfox/webservice-prototype/models"
)

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createPerson": &graphql.Field{
				Type:        IDObject,
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
					model := models.Person{
						FirstName: params.Args["firstName"].(string),
						LastName:  params.Args["lastName"].(string),
					}

					mgr := mongodb.FromContext(params.Context)
					id, err := mgr.CreatePerson(model)
					if err != nil {
						return nil, err
					}
					return id, err
				},
			},
			"deletePerson": &graphql.Field {
				Type: EmptyObject,	
				Description: "Delete an existing Person resource",
				Args: graphql.FieldConfigArgument {
					"id": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(params.Context)
					err := mgr.Delete(params.Args["id"].(string))
					return struct{}{}, err
				}
			},
		},
	})
)
