package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (context.Context, *mongo.Database) {

	host := os.Getenv("DB2_HOST")
	port := os.Getenv("DB2_PORT")
	user := os.Getenv("DB2_USER")
	pass := os.Getenv("DB2_PASS")
	name := os.Getenv("DB2_NAME")

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return ctx, client.Database(name)
}
