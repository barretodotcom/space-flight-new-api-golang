package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() *mongo.Database {

	uri :=
		`mongodb://` + os.Getenv("MONGO_NAME") + `:` + os.Getenv("MONGO_PASSWORD") + `@db:` + os.Getenv("MONGO_PORT")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}
	return client.Database("articles")
}
