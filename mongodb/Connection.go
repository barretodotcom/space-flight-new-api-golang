package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb+srv://barreto:%s@cluster0.tjwgk.mongodb.net/?retryWrites=true&w=majority", "3LEdHpxYSQmB6qJ4")
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client.Database("articles"), nil
}
