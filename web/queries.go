package web

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/webservice-prototype/db/mongodb"
)

var (
	RootQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"people": &graphql.Field{
				Type: graphql.NewList(PersonObject),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					mgr, ok := mongodb.FromContext(p.Context)
					if !ok {
						return nil, errors.New("Could not fetch PersonResourceManager from context")
					}
					people, err := mgr.Get()
					return people, err
				},
			}, // end people field
		}, // end Fields
	}) // ends object
) // end var
