package artc_http_func

import (
	"articles/mongodb"
	"articles/structs"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func PutArticle(w http.ResponseWriter, r *http.Request) {

	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())

	var article structs.Article
	json.NewDecoder(r.Body).Decode(&article)

	mongoDocument := db.Collection("articles").FindOne(context.TODO(), bson.D{{Key: "id", Value: article.Id}})

	if mongoDocument.Err() != nil {
		err := structs.Error{
			Status:  400,
			Message: "Este usuário não está cadastrado.",
		}
		errJsonBytes, _ := json.Marshal(err)

		w.WriteHeader(int(err.Status))
		w.Write((errJsonBytes))
		return
	}
	_, err := db.Collection("articles").UpdateOne(context.TODO(),
		bson.D{{Key: "id", Value: article.Id}},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "title", Value: article.Title},
				{Key: "url", Value: article.Url},
				{Key: "imageUrl", Value: article.ImageUrl},
				{Key: "newsSite", Value: article.NewsSite},
				{Key: "summary", Value: article.Summary},
				{Key: "updatedAt", Value: article.UpdatedAt},
				{Key: "featured", Value: article.Featured},
				{Key: "launches", Value: article.Launches},
				{Key: "events", Value: article.UpdatedAt}}},
		},
	)

	if err != nil {
		fmt.Println(err)
		err := structs.Error{
			Status:  400,
			Message: "Este usuário não está cadastrado.",
		}
		errJsonBytes, _ := json.Marshal(err)

		w.WriteHeader(int(err.Status))
		w.Write((errJsonBytes))
		return
	}
	var sucess structs.CreateResult = structs.CreateResult{
		Error: structs.Error{
			Status:  201,
			Message: "Documento atualizado com sucesso!",
		},
		Article: article,
	}
	sucessJson, err := json.Marshal(sucess)

	w.WriteHeader(200)
	w.Write(sucessJson)
}
