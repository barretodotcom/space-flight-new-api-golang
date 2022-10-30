package artc_http_func

import (
	"articles/mongodb"
	"articles/structs"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertArticle(w http.ResponseWriter, r *http.Request) {

	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())
	var article structs.Article

	json.NewDecoder(r.Body).Decode(&article)

	registerExists := db.Collection("articles").FindOne(r.Context(), bson.D{
		{Key: "$or", Value: []bson.M{
			{"id": article.Id},
			{"title": article.Title},
		},
		},
	})

	if registerExists.Decode(article).Error() != "mongo: no documents in result" {
		err := structs.Error{
			Status:  402,
			Message: "Já existe um registro de documento com esse título e id.",
		}
		errJson, _ := json.Marshal(err)

		w.WriteHeader(402)
		w.Write(errJson)
		return
	}

	db.Collection("articles").InsertOne(r.Context(), article)

	sucess := structs.CreateResult{
		Article: article,
		Error: structs.Error{
			Status:  201,
			Message: "Registro criado com sucesso.",
		},
	}

	sucessJson, _ := json.Marshal(sucess)
	w.WriteHeader(int(sucess.Error.Status))
	w.Write(sucessJson)
}
