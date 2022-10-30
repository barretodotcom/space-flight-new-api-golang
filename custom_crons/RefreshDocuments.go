package custom_crons

import (
	"articles/mongodb"
	"articles/structs"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RefreshDocuments() {
	fmt.Println("[ Space Flight ] Verificando necessidade de atualização")
	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())
	apiUrl := os.Getenv("ARTICLES_API_URL")

	var articles []structs.Article

	var newArticles []interface{}

	response, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println(err)
	}

	json.NewDecoder(response.Body).Decode(&articles)

	for _, v := range articles {
		documentExists := db.Collection("articles").FindOne(context.Background(), bson.D{{Key: "id", Value: v.Id}})

		if documentExists.Err() == mongo.ErrNoDocuments {
			fmt.Printf("Nenhum documento, inserindo novos registros %d \n", v.Id)

			newArticles = append(newArticles, v)
		}
	}
	if len(newArticles) > 0 {
		_, err := db.Collection("articles").InsertMany(context.TODO(), newArticles)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	db.Collection("articles").DeleteMany(context.TODO(), bson.D{{Key: "id", Value: 0}})
}
