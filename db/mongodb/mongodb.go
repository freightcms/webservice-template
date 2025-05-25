package mongodb

import (
	"fmt"
	"slices"

	"github.com/freightcms/webservice-template/db"
	"github.com/freightcms/webservice-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type resourceManager struct {
	collectionName string
	dbName         string
	session        mongo.SessionContext
}

// Get implements db.PersonResourceManager.
func (r *resourceManager) Get(query *db.PeopleQuery) ([]*models.Person, int64, error) {
	projection := bson.D{}

	// see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
	for _, fieldName := range query.Fields {
		// for security reasons we only want people to be able to query the objects that they should be able to
		if slices.Contains([]string{"id", "firstName", "lastName"}, fieldName) {
			projection = append(projection, bson.E{
				Key:   fieldName,
				Value: 1,
			})
		}
	}
	if len(query.SortBy) != 0 {
		if !slices.Contains([]string{"_id", "id"}, query.SortBy) {
			return nil, 0, fmt.Errorf("%s is not a valid sortBy option", query.SortBy)
		}
	}
	sort := bson.D{bson.E{Key: query.SortBy, Value: 1}}
	opts := options.Find().
		SetSort(sort).
		SetLimit(int64(query.PageSize)).
		SetSkip(int64((query.Page) * query.PageSize)).
		SetProjection(projection)

	cursor, err := r.collection().Find(r.session, bson.D{}, opts)
	if err != nil {
		return nil, 0, err
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

	countOpts := options.Count().
		SetLimit(int64(query.PageSize)).
		SetSkip(int64(query.PageSize) * int64(query.Page))
	count, err := r.collection().CountDocuments(r.session, bson.D{}, countOpts)
	if err != nil {
		return nil, 0, err
	}

	return results, count, nil

}

// CreatePerson implements db.PersonResourceManager.
func (r *resourceManager) CreatePerson(person models.Person) (any, error) {
	insertedResult, err := r.collection().InsertOne(r.session,
		&person,
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	return insertedResult.InsertedID, nil
}

// DeletePerson implements db.PersonResourceManager.
func (r *resourceManager) DeletePerson(id any) error {
	_, err := r.collection().DeleteOne(r.session, bson.M{"_id": id})
	return err
}

// GetById implements db.PersonResourceManager.
func (r *resourceManager) GetById(id any) (*models.Person, error) {
	var result models.Person
	filter := bson.M{"_id": id}
	if err := r.collection().FindOne(r.session, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePerson implements db.PersonResourceManager.
func (r *resourceManager) UpdatePerson(id any, person models.Person) error {
	result, err := r.collection().UpdateOne(r.session, bson.M{"_id": id}, person)

	if result.MatchedCount == 0 {
		return fmt.Errorf("could not find Person with id %s", id)
	}
	return err
}

func NewPersonManager(dbName, collectionName string, session mongo.SessionContext) db.PersonResourceManager {
	return &resourceManager{
		dbName:         dbName,
		collectionName: collectionName,
		session:        session,
	}
}

func (r *resourceManager) collection() *mongo.Collection {
	coll := r.session.Client().Database("freightcms").Collection("people")
	return coll
}
