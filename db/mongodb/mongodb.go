package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/squishedfox/webservice-prototype/db"
	"github.com/squishedfox/webservice-prototype/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PersonResourceManagerContextKey string

const (
	// ContextKey used to fetch or put the Person Resource Manager into the context
	ContextKey PersonResourceManagerContextKey = "personResourceManagerContextKey"
)

type resourceManager struct {
	session mongo.SessionContext
}

// Get implements db.PersonResourceManager.
func (r *resourceManager) Get() ([]*models.Person, error) {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	cursor, err := coll.Find(r.session, bson.D{}, nil)
	if err != nil {
		return nil, err
	}
	results := []*models.Person{}
	for cursor.Next(r.session) {
		var result models.Person
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Error occured fetching record %s\n", err.Error())
			continue
		}
		results = append(results, &result)
	}
	return results, nil

}

// WithContext fetches the mongo db session context from that passed argument (parent context)
// ,appends the person manager and returns all with the new context.
func WithContext(session mongo.SessionContext) context.Context {
	if session == nil {
		panic("Could not fetch session from context")
	}
	mgr := NewPersonManager(session)
	return context.WithValue(session, ContextKey, mgr)
}

// FromContext gets the Resource Manager from the context passsed.
func FromContext(ctx context.Context) db.PersonResourceManager {
	val := ctx.Value(ContextKey)
	if val == nil {
		panic(errors.New("could not fetch PersonResourceManager from context"))
	}

	return val.(*resourceManager)
}

// CreatePerson implements db.PersonResourceManager.
func (r *resourceManager) CreatePerson(person models.Person) (interface{}, error) {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	insertedResult, err := coll.InsertOne(r.session,
		&person,
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	return insertedResult.InsertedID, nil
}

// DeletePerson implements db.PersonResourceManager.
func (r *resourceManager) DeletePerson(id interface{}) error {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	_, err := coll.DeleteOne(r.session, bson.M{"_id": id})
	return err
}

// GetById implements db.PersonResourceManager.
func (r *resourceManager) GetById(id interface{}) (*models.Person, error) {
	var result models.Person
	filter := bson.M{"_id": id}
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	if err := coll.FindOne(r.session, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePerson implements db.PersonResourceManager.
func (r *resourceManager) UpdatePerson(id interface{}, person models.Person) error {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	result, err := coll.UpdateOne(r.session, bson.M{"_id": id}, person)

	if result.MatchedCount == 0 {
		return errors.New(fmt.Sprintf("Could not find Person with id %s", id))
	}
	return err
}

func NewPersonManager(session mongo.SessionContext) db.PersonResourceManager {
	return &resourceManager{session: session}
}
