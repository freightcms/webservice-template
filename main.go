package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/freightcms/webservice-template/db"
	"github.com/freightcms/webservice-template/db/mongodb"
	"github.com/freightcms/webservice-template/web"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// addMongoDbMiddleware adds the CarrierResourceManager to the echo context so that it can be
// be recovered from the db.DbContext object
func addMongoDbMiddleware(client *mongo.Client, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := client.StartSession()
		if err != nil {
			return err
		}
		requestContext := c.Request().Context()
		defer session.EndSession(requestContext)

		sessionContext := mongo.NewSessionContext(requestContext, session)
		dbContext := db.DbContext{
			Context:                requestContext,
			CarrierResourceManager: mongodb.NewCarrierManager(sessionContext),
		}
		wrappedContext := web.AppContext{
			Context:   c,
			DbContext: dbContext,
		}
		return next(wrappedContext)
	}
}

var (
	port           int
	host           string
	dbName         string
	collectionName string
)

func main() {

	flag.IntVar(&port, "p", 8080, "Port to run application on")
	flag.StringVar(&host, "h", "0.0.0.0", "Host address to run application on")
    flag.StringVar(&dbName "database", "Name of the database to use when connecting")
    flag.StringVar(&cocollectionName "collection", "Name of the collection in mongodb to use when connecting")
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
	fmt.Println("Pinging server...")
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		return
	}

    if dbName == "" {
        dbName = os.Getenv("DATABASE_NAME")
        if dbName == "" {
            log.Fatal("Could not get database name from environment or cli option '--database=...'")
        }
    }

    if collectionName == "" {
        dbName = os.Getenv("COLLECTION_NAME")
        if dbName == "" {
            log.Fatal("Could not get collection name from environment or cli option '--collection=...'")
        }
    }

	fmt.Println("Done")
	fmt.Println("Setting up handlers and routes")

	server := echo.New()
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return addMongoDbMiddleware(client, next)
	})

	web.Register(server)

	fmt.Println("Done")
	hostname := fmt.Sprintf("%v:%d", host, port)
	fmt.Printf("Start server at %s\n", hostname)
	http.ListenAndServe(hostname, server)
}
