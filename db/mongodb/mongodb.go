package mongodb

import (
	"github.com/squishedfox/webservice-prototype/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type resourceManager struct {
	session mongo.SessionContext
}

func NewPersonManager(session mongo.SessionContext) db.PersonResourceManager {
	return &resourceManager{session: session}
}
