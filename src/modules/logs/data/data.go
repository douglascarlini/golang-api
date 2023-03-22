package data

import (
	"context"
	database "src/shared/database/mongo"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogData struct {
	DB    *mongo.Database
	CTX   context.Context
	TABLE string
}

func NewLogData() *LogData {
	ctx, db := database.Connect()
	return &LogData{
		TABLE: "logs",
		CTX:   ctx,
		DB:    db,
	}
}
