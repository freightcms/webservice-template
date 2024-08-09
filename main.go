package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting application...")
	if err := dotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_SERVER")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Disconnect(context.Background())
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
		return
	}

	personObject := graphql.NewObject(graphql.ObjectConfig{
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
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"people": &graphql.Field{
				Type: graphql.NewList(personObject),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					session, err := client.StartSession()
					if err != nil {
						return nil, err
					}
					defer session.EndSession(p.Context)
					coll := client.Database("graphql_mongo_prototype").Collection("people")
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
	rootSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	h := handler.New(&handler.Config{
		Schema:   &rootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	server := http.NewServeMux()
	server.Handle("/graphql", h)
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
		w.Header().Set("ContentType", "applicatoin/json")
	})

	http.ListenAndServe(":8080", server)
}
