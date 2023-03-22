package data

import (
	"src/modules/logs/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *LogData) Add(item *models.Log) (string, error) {

	var err error
	var res *mongo.InsertOneResult

	item.Created = time.Now()
	collection := r.DB.Collection(r.TABLE)

	if res, err = collection.InsertOne(r.CTX, &item); err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil

}
