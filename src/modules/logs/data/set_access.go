package data

import (
	"src/modules/logs/models"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LogData) SetAccess(item *models.Log, id string) error {

	var err error
	var _id primitive.ObjectID

	if _id, err = primitive.ObjectIDFromHex(id); err != nil {
		return err
	}

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"user_id": item.UserID, "app_id": item.AppID}}

	if _, err = r.DB.Collection(r.TABLE).UpdateOne(r.CTX, filter, update); err != nil {
		return err
	}

	return nil

}
