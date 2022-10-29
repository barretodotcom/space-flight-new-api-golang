package artc_http_func

import (
	"articles/mongodb"
	"articles/structs"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())

	array, err := db.Collection("articles").Find(context.TODO(), bson.D{})

	var articles structs.Articles

	for array.Next(r.Context()) {
		result := structs.Article{}

		err := array.Decode(&result)
		if err != nil {
			log.Panic(err.Error())
		}
		articles.Articles = append(articles.Articles, result)
	}

	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}

	jsonBytes, err := json.Marshal(articles)

	if err != nil {
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}

	w.Write(jsonBytes)
}
