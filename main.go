package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/graphql-go/handler"
	"github.com/squishedfox/webservice-prototype/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	fmt.Println("Starting application...")
	if err := dotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_SERVER")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		return
	}

	h := handler.New(&handler.Config{
		Schema:   &web.RootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	server := http.NewServeMux()
	server.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		session, err := client.StartSession()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Header().Set("ContentType", "application/json")
			return
		}
		defer session.EndSession(r.Context())
		h.ServeHTTP(w, r.WithContext(mongo.NewSessionContext(r.Context(), session)))
	})

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
		w.Header().Set("ContentType", "application/json")
	})

	http.ListenAndServe(":8080", server)
}
