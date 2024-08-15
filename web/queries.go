package web

import (
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
					mgr := mongodb.FromContext(p.Context)
					people, err := mgr.Get()
					return people, err
				},
			}, // end people field
		}, // end Fields
	}) // ends object
) // end var
