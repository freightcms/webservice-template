package web

import (
	"github.com/graphql-go/graphql"
)

var (
	EmptyObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig({
		Name: "Empty",
		Fields: graphql.Fields{},
	}),
	IDObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig({
		Name: "ID",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Name: "id",
				Type graphql.String,
			},
		},
	}),
	PersonObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
)
